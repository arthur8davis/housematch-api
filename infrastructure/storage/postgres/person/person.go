package person

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.persons"
)

var (
	_psqlGetById = `SELECT p.id,
				       p.document_type,
				       p.document,
				       p.names,
				       p.lastname,
				       p.m_lastname,
				       p.phone,
				       p.gender,
				       p.marital_status,
				       p.date_birth,
				       lp.id,
				       lp.country,
				       lp.city,
				       lp.province,
				       lp.district
				FROM domain.persons p INNER JOIN domain.locations lp ON p.location_id = lp.id
				WHERE p.id = $1;`
	_psqlGetAll = `SELECT p.id,
				       p.document_type,
				       p.document,
				       p.names,
				       p.lastname,
				       p.m_lastname,
				       p.phone,
				       p.gender,
				       p.marital_status,
				       p.date_birth,
				       lp.id,
				       lp.country,
				       lp.city,
				       lp.province,
				       lp.district
				FROM domain.persons p INNER JOIN domain.locations lp ON p.location_id = lp.id;`
	_psqlInsert = `INSERT INTO domain.persons (id, document_type, document, names, lastname, m_lastname, phone, gender, marital_status, date_birth, photo, location_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
	_psqlUpdate = `UPDATE domain.persons SET "document_type"=$2, "document"=$3, "names"=$4, "lastname"=$5, "m_lastname"=$6, "phone"=$7, "gender"=$8, "marital_status"=$9, "address"=$10, "date_birth"=$11, "location_id"=$12 WHERE id=$1;`
	_psqlDelete = `DELETE FROM domain.persons WHERE id=$1;`
)

type Person struct {
	db *sql.DB
}

func New(db *sql.DB) Person {
	return Person{db}
}

func (p Person) GetByIdStorage(id uuid.UUID) (*model.PersonSecondLevel, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := p.scanRowWithLocation(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (p Person) GetAllStorage() (model.PersonsOutput, error) {
	stmt, err := p.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.PersonsOutput
	var m model.PersonSecondLevel

	for rows.Next() {
		m, err = p.scanRowWithLocation(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Person) CreateStorage(person model.Person) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	person.ID = newId

	args := p.readModelPerson(person)

	stmt, err := p.db.Prepare(_psqlInsert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return &newId, nil
}

func (p Person) UpdateStorage(id uuid.UUID, person model.Person) (bool, error) {
	person.ID = id

	args := p.readModelPerson(person)

	stmt, err := p.db.Prepare(_psqlUpdate)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("nothing rows updated, error: %v", err)
	}

	return true, nil
}

func (p Person) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlDelete)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("nothing rows delete, error: %v", err)
	}

	return true, nil
}

func (p Person) readModelPerson(person model.Person) []any {
	v := reflect.ValueOf(person)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (p Person) scanRow(s pgx.Row) (model.Person, error) {
	m := model.Person{}

	err := s.Scan(
		&m.ID,
		&m.DocumentType,
		&m.Document,
		&m.Names,
		&m.Lastname,
		&m.MLastname,
		&m.Phone,
		&m.Gender,
		&m.MaritalStatus,
		&m.DateBirth,
		&m.LocationID,
	)
	if err != nil {
		return model.Person{}, err
	}

	return m, nil
}

func (p Person) scanRowWithLocation(s pgx.Row) (model.PersonSecondLevel, error) {
	m := model.PersonSecondLevel{}

	err := s.Scan(
		&m.ID,
		&m.DocumentType,
		&m.Document,
		&m.Names,
		&m.Lastname,
		&m.MLastname,
		&m.Phone,
		&m.Gender,
		&m.MaritalStatus,
		&m.DateBirth,
		&m.LocationPerson.ID,
		&m.LocationPerson.Country,
		&m.LocationPerson.City,
		&m.LocationPerson.Province,
		&m.LocationPerson.District,
	)
	if err != nil {
		return model.PersonSecondLevel{}, err
	}

	return m, nil
}
