package user

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.users"
)

var (
	_psqlGetById = `SELECT u.id,
						u.user,
						u.password,
						u.email,
						u.theme,
						p.id,
						p.document_type,
						p.document,
						p.names,
						p.lastname,
						p.m_lastname,
						p.phone,
						p.gender,
						p.marital_status,
						p.date_birth,
						p.location_id
					FROM domain.users u INNER JOIN domain.persons p ON p.id = u.person_id
					WHERE u.id = $1`
	_psqlGetByUsernameOrEmail = `SELECT id,
									user,
									password,
									email,
									theme,
									person_id
								From domain.users
								WHERE "user" = $1 OR email = $2`
	_psqlGetAll = `SELECT u.id,
				       u.user,
				       u.password,
				       u.email,
				       u.theme,
				       p.id,
				       p.document_type,
				       p.document,
				       p.names,
				       p.lastname,
				       p.m_lastname,
				       p.phone,
				       p.gender,
				       p.marital_status,
				       p.date_birth,
				       p.location_id
				FROM domain.users u INNER JOIN domain.persons p ON p.id = u.person_id`
	_psqlGetAllWithRoles = `SELECT u.id,
								   u.user,
								   u.password,
								   u.email,
								   u.theme,
								   r.id,
								   r.name,
								   r.description,
								   r."order"
							FROM domain.users_roles ur INNER JOIN domain.users u ON u.id = ur.user_id
							INNER JOIN domain.roles r ON ur.role_id = r.id`
	_psqlInsert = `INSERT INTO domain.users (id, "user", "password", "email", "theme", "person_id") VALUES ($1, $2, $3, $4, $5, $6)`
	_psqlUpdate = `UPDATE domain.users SET "user"=$2, "password"=$3, "email"=$4, "theme"=$5, "person_id"=$6 WHERE id=$1`
	_psqlDelete = `DELETE FROM domain.users WHERE id=$1`
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) GetByIdStorage(id uuid.UUID) (*model.UserSecondLevel, error) {
	args := []any{id}

	stmt, err := u.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := u.scanRowWithPerson(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (u User) GetByUsernameOrEmailStorage(username, email string) (*model.User, error) {
	args := []any{username, email}

	stmt, err := u.db.Prepare(_psqlGetByUsernameOrEmail)
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

func (u User) GetAllStorage() (model.UsersSecondLevel, error) {
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

	var ms model.UsersSecondLevel
	var m model.UserSecondLevel

	for rows.Next() {
		m, err = u.scanRowWithPerson(rows)
		if err != nil {
			break
		}
		fmt.Println("model", m)
		ms = append(ms, m)
	}
	fmt.Println("models", ms)

	return ms, nil
}

func (u User) GetAllWithRolesStorage() (model.UsersWithRolesOutput, error) {
	stmt, err := u.db.Prepare(_psqlGetAllWithRoles)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.UsersWithRoles
	var m model.UserWithRole

	for rows.Next() {
		m, err = u.scanRowWithRole(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return makeUsersWithRoles(ms), nil
}

func (u User) CreateStorage(user model.User) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	user.ID = newId

	args := u.readModelUser(user)

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

func (u User) UpdateStorage(id uuid.UUID, user model.User) (bool, error) {
	user.ID = id

	args := u.readModelUser(user)

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

func (u User) DeleteStorage(id uuid.UUID) (bool, error) {
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

func (u User) readModelUser(user model.User) []any {
	v := reflect.ValueOf(user)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (u User) scanRow(s pgx.Row) (model.User, error) {
	m := model.User{}

	err := s.Scan(
		&m.ID,
		&m.User,
		&m.Password,
		&m.Email,
		&m.Theme,
		&m.PersonID,
	)
	if err != nil {
		return model.User{}, err
	}

	return m, nil
}

func (u User) scanRowWithPerson(s pgx.Row) (model.UserSecondLevel, error) {
	m := model.UserSecondLevel{}

	err := s.Scan(
		&m.ID,
		&m.User,
		&m.Password,
		&m.Email,
		&m.Theme,
		&m.Person.ID,
		&m.Person.DocumentType,
		&m.Person.Document,
		&m.Person.Names,
		&m.Person.Lastname,
		&m.Person.MLastname,
		&m.Person.Phone,
		&m.Person.Gender,
		&m.Person.MaritalStatus,
		&m.Person.DateBirth,
		&m.Person.LocationID,
	)
	if err != nil {
		return model.UserSecondLevel{}, err
	}

	return m, nil
}

func (u User) scanRowWithRole(s pgx.Row) (model.UserWithRole, error) {
	m := model.UserWithRole{}

	err := s.Scan(
		&m.ID,
		&m.User,
		&m.Password,
		&m.Email,
		&m.Theme,
		&m.Role.ID,
		&m.Role.Name,
		&m.Role.Description,
		&m.Role.Order,
	)
	if err != nil {
		return model.UserWithRole{}, err
	}

	return m, nil
}

func makeUsersWithRoles(ms model.UsersWithRoles) model.UsersWithRolesOutput {
	result := ms.GetUserWithRole()

	return result
}
