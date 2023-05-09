package view

import (
	"database/sql"
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.views"
)

var (
	_psqlGetById = `select v.id,
    	v.name,
    	v.description,
    	v.url,
    	v.icon,
    	m.id,
    	m.name,
    	m.description,
    	m.icon,
    	m."order"
		from domain.views v inner join domain.modules m on m.id = v.module_id WHERE v.id = $1;`
	_psqlGetAll = `SELECT v.id,
    	v.name,
    	v.description,
    	v.url,
    	v.icon,
    	m.id,
    	m.name,
    	m.description,
    	m.icon,
    	m."order"
		FROM domain.views v INNER JOIN domain.modules m ON m.id = v.module_id;`
	_psqlInsert = `INSERT INTO domain.views (id, "module_id", "name", "description", "url", "icon") VALUES ($1, $2, $3, $4, $5, $6)`
	_psqlUpdate = `UPDATE domain.views SET "module_id"=$2, "name"=$3, "description"=$4, "url"=$5, "icon"=$6 WHERE id=$1`
	_psqlDelete = `DELETE FROM domain.views WHERE id=$1`
)

type View struct {
	db *sql.DB
}

func New(db *sql.DB) View {
	return View{db}
}

func (v View) GetStorageById(id uuid.UUID) (*model.ViewOutput, error) {
	args := []any{id}

	stmt, err := v.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := v.scanRowWithModule(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (v View) GetStorageAll() (model.ViewsOutput, error) {
	stmt, err := v.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.ViewsOutput
	var m model.ViewOutput

	for rows.Next() {
		m, err = v.scanRowWithModule(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (v View) CreateStorage(view model.View) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	view.ID = newId

	args := v.readModelView(view)

	stmt, err := v.db.Prepare(_psqlInsert)
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

func (v View) UpdateStorage(id uuid.UUID, view model.View) (bool, error) {
	view.ID = id

	args := v.readModelView(view)

	stmt, err := v.db.Prepare(_psqlUpdate)
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

func (v View) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := v.db.Prepare(_psqlDelete)
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

func (v View) readModelView(view model.View) []any {
	r := reflect.ValueOf(view)
	values := make([]interface{}, r.NumField())
	for i := 0; i < r.NumField(); i++ {
		values[i] = r.Field(i).Interface()
	}

	return values
}

func (v View) scanRow(s pgx.Row) (model.View, error) {
	m := model.View{}

	err := s.Scan(
		&m.ID,
		&m.ModuleID,
		&m.Name,
		&m.Description,
		&m.URL,
		&m.Icon,
	)
	if err != nil {
		return model.View{}, err
	}

	return m, nil
}

func (v View) scanRowWithModule(s pgx.Row) (model.ViewOutput, error) {
	m := model.ViewOutput{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Description,
		&m.URL,
		&m.Icon,
		&m.Module.ID,
		&m.Module.Name,
		&m.Module.Description,
		&m.Module.Icon,
		&m.Module.Order,
	)
	if err != nil {
		return model.ViewOutput{}, err
	}

	return m, nil
}
