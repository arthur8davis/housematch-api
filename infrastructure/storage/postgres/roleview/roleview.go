package roleview

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.roles_views"
)

var (
	_psqlGetById = `
		select rv.view_order,
		       rv.view_position,
		       r.id,
		       r.name,
		       r.description,
		       r."order",
		       v.id,
		       v.module_id,
		       v.name,
		       v.description,
		       v.url,
		       v.icon
		from domain.roles_views rv
		         inner join domain.roles r on r.id = rv.role_id
		         inner join domain.views v on v.id = rv.view_id
		where rv.role_id = $1 and rv.view_id = $2;`
	_psqlGetAll = `
		select rv.view_order,
		       rv.view_position,
		       r.id,
		       r.name,
		       r.description,
		       r."order",
		       v.id,
		       v.module_id,
		       v.name,
		       v.description,
		       v.url,
		       v.icon
		from domain.roles_views rv
		         inner join domain.roles r on r.id = rv.role_id
		         inner join domain.views v on v.id = rv.view_id;`
	_psqlInsert = `INSERT INTO domain.roles_views ("role_id", "view_id", "view_order", "view_position") VALUES ($1, $2, $3, $4);`
	_psqlUpdate = `UPDATE domain.roles_views SET "view_order"=$3, "view_position"=$4 WHERE "role_id"=$1 and "view_id"=$2;`
	_psqlDelete = `DELETE FROM domain.roles_views WHERE "role_id"=$1 and "view_id"=$2;`
)

type RoleView struct {
	db *sql.DB
}

func New(db *sql.DB) RoleView {
	return RoleView{db}
}

func (v RoleView) GetByIDsStorage(roleID, viewID uuid.UUID) (*model.RoleViewOutput, error) {
	args := []any{roleID, viewID}

	stmt, err := v.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := v.scanRowRoleView(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (v RoleView) GetAllStorage() (model.RoleViewOutputs, error) {
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

	var ms model.RoleViewOutputs
	var m model.RoleViewOutput

	for rows.Next() {
		m, err = v.scanRowRoleView(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (v RoleView) AssignmentStorage(roleView model.RoleView) (bool, error) {
	args := v.readModelRoleView(roleView)

	stmt, err := v.db.Prepare(_psqlInsert)
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

func (v RoleView) UpdateStorage(roleView model.RoleView) (bool, error) {
	args := v.readModelRoleView(roleView)

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

func (v RoleView) DeleteStorage(roleID, viewID uuid.UUID) (bool, error) {
	args := []any{roleID, viewID}

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

func (v RoleView) readModelRoleView(view model.RoleView) []any {
	r := reflect.ValueOf(view)
	values := make([]interface{}, r.NumField())
	for i := 0; i < r.NumField(); i++ {
		values[i] = r.Field(i).Interface()
	}

	return values
}

func (v RoleView) scanRow(s pgx.Row) (model.RoleView, error) {
	m := model.RoleView{}

	err := s.Scan(
		&m.RoleID,
		&m.ViewID,
		&m.ViewOrder,
		&m.ViewPosition,
	)
	if err != nil {
		return model.RoleView{}, err
	}

	return m, nil
}

func (v RoleView) scanRowRoleView(s pgx.Row) (model.RoleViewOutput, error) {
	m := model.RoleViewOutput{}

	err := s.Scan(
		&m.ViewOrder,
		&m.ViewPosition,
		&m.Role.ID,
		&m.Role.Name,
		&m.Role.Description,
		&m.Role.Order,
		&m.View.ID,
		&m.View.ModuleID,
		&m.View.Name,
		&m.View.Description,
		&m.View.URL,
		&m.View.Icon,
	)
	if err != nil {
		return model.RoleViewOutput{}, err
	}

	return m, nil
}
