package media

import (
	"database/sql"
	"fmt"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

var (
	_psqlInsert = `INSERT INTO domain.medias (id, "name", "url", "size", "type") VALUES ($1, $2, $3, $4, $5)`
	_psqlDelete = `DELETE FROM domain.medias WHERE id=$1`
)

type Media struct {
	db *sql.DB
}

func New(db *sql.DB) Media {
	return Media{db}
}

func (m Media) CreateStorage(media model.Media) error {
	args := m.readModelMedia(media)

	stmt, err := m.db.Prepare(_psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func (m Media) readModelMedia(module model.Media) []any {
	v := reflect.ValueOf(module)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (m Media) DeleteStorage(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := m.db.Prepare(_psqlDelete)
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

func (m Media) scanRow(s pgx.Row) (model.Media, error) {
	md := model.Media{}

	err := s.Scan(
		&md.ID,
		&md.Name,
		&md.URL,
		&md.Size,
		&md.Type,
	)
	if err != nil {
		return model.Media{}, err
	}

	return md, nil
}
