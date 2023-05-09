package personlocation

import (
	"database/sql"
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

var (
	_psqlInsertPerson   = `INSERT INTO domain.persons (id, document_type, document, names, lastname, m_lastname, phone, gender, marital_status, date_birth,photo, location_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
	_psqlInsertLocation = `INSERT INTO domain.locations (id, country, city, province, district, address, lat, long) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_psqlUpdatePerson   = `UPDATE domain.persons SET "document_type"=$2, "document"=$3, "names"=$4, "lastname"=$5, "m_lastname"=$6, "phone"=$7, "gender"=$8, "marital_status"=$9, "date_birth"=$10,"photo"=$11, "location_id"=$12 WHERE id=$1;`
	_psqlUpdateLocation = `UPDATE domain.locations SET "country"=$2, "city"=$3, "province"=$4, "district"=$5, "address"=$6, "lat"=$7, "long"=$8 WHERE id=$1`
)

type PersonLocation struct {
	db *sql.DB
}

func New(db *sql.DB) PersonLocation {
	return PersonLocation{db}
}

func (p PersonLocation) CreateStorage(personLocation model.PersonLocation) (*uuid.UUID, error) {

	idLocation, err := p.CreateLocation(personLocation.Location)

	if err != nil {
		return nil, err
	}

	idPerson, err := p.CreatePerson(personLocation.Person, idLocation)

	if err != nil {
		return nil, err
	}

	return idPerson, nil
}

func (p PersonLocation) CreateLocation(location model.Location) (*uuid.UUID, error) {
	newIdLocation, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}

	location.ID = newIdLocation
	args := p.readModelLocation(location)

	stmt, err := p.db.Prepare(_psqlInsertLocation)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return &newIdLocation, nil
}

func (p PersonLocation) CreatePerson(person model.Person, idLocation *uuid.UUID) (*uuid.UUID, error) {
	newIdPerson, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}

	person.ID = newIdPerson
	person.LocationID = *idLocation
	args := p.readModelPerson(person)

	fmt.Println(args)

	stmt, err := p.db.Prepare(_psqlInsertPerson)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return &newIdPerson, nil
}

func (p PersonLocation) UpdateStorage(personID uuid.UUID, personLocation model.PersonLocation) (bool, error) {

	_, err := p.UpdatePerson(personID, personLocation.Person)
	if err != nil {
		return false, err
	}

	_, err = p.UpdateLocation(personLocation.Person.LocationID, personLocation.Location)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (p PersonLocation) UpdateLocation(id uuid.UUID, Location model.Location) (bool, error) {
	Location.ID = id

	args := p.readModelLocation(Location)

	stmt, err := p.db.Prepare(_psqlUpdateLocation)
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

func (p PersonLocation) UpdatePerson(id uuid.UUID, Person model.Person) (bool, error) {
	Person.ID = id

	args := p.readModelPerson(Person)

	stmt, err := p.db.Prepare(_psqlUpdatePerson)
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

func (p PersonLocation) readModelLocation(location model.Location) []any {
	v := reflect.ValueOf(location)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (p PersonLocation) readModelPerson(person model.Person) []any {
	v := reflect.ValueOf(person)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return values
}

func (p PersonLocation) scanRow(s pgx.Row) (model.Person, error) {
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

func (p PersonLocation) scanRowWithLocation(s pgx.Row) (model.PersonSecondLevel, error) {
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
