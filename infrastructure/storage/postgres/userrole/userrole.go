package userrole

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.users_roles"
)

var (
	_psqlGetByIds = `
		SELECT u.id,
    			u."user",
    			u.password,
    			u.email,
    			u.theme,
    			r.id,
    			r.name,
    			r.description,
    			r."order"
		FROM domain.users_roles ur
		         INNER JOIN domain.roles r on r.id = ur.role_id
		         INNER JOIN domain.users u on u.id = ur.user_id
		WHERE ur.user_id = $1 AND ur.role_id = $2;`
	_psqlGetAll = `
		SELECT u.id,
    			u."user",
    			u.password,
    			u.email,
    			u.theme,
    			r.id,
    			r.name,
    			r.description,
    			r."order"
		FROM domain.users_roles ur
		         INNER JOIN domain.roles r on r.id = ur.role_id
		         INNER JOIN domain.users u on u.id = ur.user_id
		WHERE ur.user_id = $1;`
	_psqlInsert = `INSERT INTO domain.users_roles ("user_id", "role_id") VALUES ($1, $2);`
	_psqlDelete = `DELETE FROM domain.users_roles WHERE "user_id"=$1 and "role_id"=$2;`
)

type UserRole struct {
	db *sql.DB
}

func New(db *sql.DB) UserRole {
	return UserRole{db}
}

func (ur UserRole) GetByIDsStorage(userID, roleID uuid.UUID) (*model.UserRoleOutput, error) {
	fmt.Println("#######################x")
	args := []any{userID, roleID}

	stmt, err := ur.db.Prepare(_psqlGetByIds)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := ur.scanRowUserRole(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}
	fmt.Println("#######################y")
	return &m, nil
}

func (ur UserRole) GetAllByUserIDStorage(userID uuid.UUID) (model.UserRoleOutputs, error) {
	args := []any{userID}

	stmt, err := ur.db.Prepare(_psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.UserRoleOutputs
	var m model.UserRoleOutput

	for rows.Next() {
		m, err = ur.scanRowUserRole(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (ur UserRole) AssignmentStorage(userID, roleID uuid.UUID) (bool, error) {
	args := []any{userID, roleID}

	stmt, err := ur.db.Prepare(_psqlInsert)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ur UserRole) DeleteStorage(userID, roleID uuid.UUID) (bool, error) {
	args := []any{userID, roleID}

	stmt, err := ur.db.Prepare(_psqlDelete)
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

func (ur UserRole) readModelUserRole(role model.UserRole) []any {
	v := reflect.ValueOf(role)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (ur UserRole) scanRow(s pgx.Row) (model.UserRole, error) {
	m := model.UserRole{}

	err := s.Scan(
		&m.UserID,
		&m.RoleID,
	)
	if err != nil {
		return model.UserRole{}, err
	}

	return m, nil
}

func (ur UserRole) scanRowUserRole(s pgx.Row) (model.UserRoleOutput, error) {
	m := model.UserRoleOutput{}

	err := s.Scan(
		&m.User.ID,
		&m.User.User,
		&m.User.Password,
		&m.User.Email,
		&m.User.Theme,
		&m.Role.ID,
		&m.Role.Name,
		&m.Role.Description,
		&m.Role.Order,
	)
	if err != nil {
		return model.UserRoleOutput{}, err
	}

	return m, nil
}
