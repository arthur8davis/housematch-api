package module

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.modules"
)

var (
	_psqlGetById = `SELECT * FROM domain.modules WHERE id = $1`
	_psqlGetAll  = `SELECT * FROM domain.modules`
	_psqlInsert  = `INSERT INTO domain.modules (id, "name", "description", "icon", "order") VALUES ($1, $2, $3, $4, $5)`
	_psqlUpdate  = `UPDATE domain.modules SET "name"=$2, "description"=$3, "icon"=$4, "order"=$5 WHERE id=$1`
	_psqlDelete  = `DELETE FROM domain.modules WHERE id=$1`
)

type Module struct {
	db *sql.DB
}

func New(db *sql.DB) Module {
	return Module{db}
}

func (u Module) GetStorageById(id uuid.UUID) (*model.Module, error) {
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

func (u Module) GetStorageAll() (model.Modules, error) {
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

	var ms model.Modules
	var m model.Module

	for rows.Next() {
		m, err = u.scanRow(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u Module) CreateStorage(module model.Module) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	module.ID = newId

	args := u.readModelModule(module)

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

func (u Module) UpdateStorage(id uuid.UUID, module model.Module) (bool, error) {
	module.ID = id

	args := u.readModelModule(module)

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

func (u Module) DeleteStorage(id uuid.UUID) (bool, error) {
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

func (u Module) readModelModule(module model.Module) []any {
	v := reflect.ValueOf(module)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (u Module) scanRow(s pgx.Row) (model.Module, error) {
	m := model.Module{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Description,
		&m.Icon,
		&m.Order,
	)
	if err != nil {
		return model.Module{}, err
	}

	return m, nil
}
