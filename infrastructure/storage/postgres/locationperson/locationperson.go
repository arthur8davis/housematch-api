package locationperson

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.locations"
)

var (
	_psqlGetById = `SELECT * FROM domain.locations WHERE id = $1`
	_psqlGetAll  = `SELECT * FROM domain.locations`
	_psqlInsert  = `INSERT INTO domain.locations (id, "country", "city", "province", "district", "address", "lat", "long") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_psqlUpdate  = `UPDATE domain.locations SET "country"=$2, "city"=$3, "province"=$4, "district"=$5, "address"=$6, "lat"=$7, "long"=$8 WHERE id=$1`
	_psqlDelete  = `DELETE FROM domain.locations WHERE id=$1`
)

type Location struct {
	db *sql.DB
}

func New(db *sql.DB) Location {
	return Location{db}
}

func (u Location) GetByIdStorage(id uuid.UUID) (*model.Location, error) {
	args := []any{id}

	stmt, err := u.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := u.scanRow(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (u Location) GetAllStorage() (model.Locations, error) {
	stmt, err := u.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Locations
	var m model.Location

	for rows.Next() {
		m, err = u.scanRow(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u Location) CreateStorage(locationPerson model.Location) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	locationPerson.ID = newId

	args := u.readModelLocationPerson(locationPerson)

	stmt, err := u.db.Prepare(_psqlInsert)
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

func (u Location) UpdateStorage(id uuid.UUID, locationPerson model.Location) (bool, error) {
	locationPerson.ID = id

	args := u.readModelLocationPerson(locationPerson)

	stmt, err := u.db.Prepare(_psqlUpdate)
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

func (u Location) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := u.db.Prepare(_psqlDelete)
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

func (u Location) readModelLocationPerson(locationPerson model.Location) []any {
	v := reflect.ValueOf(locationPerson)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (u Location) scanRow(s pgx.Row) (model.Location, error) {
	m := model.Location{}

	err := s.Scan(
		&m.ID,
		&m.Country,
		&m.City,
		&m.Province,
		&m.District,
		&m.Address,
		&m.Lat,
		&m.Long,
	)
	if err != nil {
		return model.Location{}, err
	}

	return m, nil
}
