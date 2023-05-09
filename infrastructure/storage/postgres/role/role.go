package role

import (
	"database/sql"
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.roles"
)

var (
	_psqlGetById = `SELECT * FROM domain.roles WHERE id = $1`
	_psqlGetAll  = `SELECT * FROM domain.roles`
	_psqlInsert  = `INSERT INTO domain.roles (id, "name", "description", "order") VALUES ($1, $2, $3, $4)`
	_psqlUpdate  = `UPDATE domain.roles SET "name"=$2, "description"=$3, "order"=$4 WHERE id=$1`
	_psqlDelete  = `DELETE FROM domain.roles WHERE id=$1`
)

type Role struct {
	db *sql.DB
}

func New(db *sql.DB) Role {
	return Role{db}
}

func (r Role) GetStorageById(id uuid.UUID) (*model.Role, error) {
	args := []any{id}

	stmt, err := r.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := r.scanRow(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r Role) GetStorageAll() (model.Roles, error) {
	stmt, err := r.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Roles
	var m model.Role

	for rows.Next() {
		m, err = r.scanRow(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (r Role) CreateStorage(role model.Role) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	role.ID = newId

	args := r.readModelRole(role)

	stmt, err := r.db.Prepare(_psqlInsert)
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

func (r Role) UpdateStorage(id uuid.UUID, role model.Role) (bool, error) {
	role.ID = id

	args := r.readModelRole(role)

	stmt, err := r.db.Prepare(_psqlUpdate)
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

func (r Role) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := r.db.Prepare(_psqlDelete)
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

func (r Role) readModelRole(role model.Role) []any {
	v := reflect.ValueOf(role)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (r Role) scanRow(s pgx.Row) (model.Role, error) {
	m := model.Role{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Description,
		&m.Order,
	)
	if err != nil {
		return model.Role{}, err
	}

	return m, nil
}
