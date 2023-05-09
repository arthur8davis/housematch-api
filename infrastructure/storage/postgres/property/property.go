package property

import (
	"database/sql"
	"fmt"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"reflect"
)

const (
	table = "domain.properties"
)

var (
	_psqlGetById = `SELECT p.id,
						   l.id,
						   l.country,
						   l.city,
						   l.province,
						   l.district,
						   l.address,
						   l.lat,
						   l.long,
						   p.description,
						   p.type,
						   p.length,
						   p.width,
						   p.area,
						   p.floor,
						   p.number_of_floors,
						   p.rooms,
						   p.bathrooms,
						   p.yard,
						   p.terrace,
						   p.living_room,
						   p.laundry_room,
						   p.kitchen,
						   p.garage,
						   u.id,
						   u.user,
						   u.password,
						   u.email,
						   u.theme,
						   u.person_id
					FROM domain.properties p
							 INNER JOIN domain.users u ON p.user_id = u.id
							 INNER JOIN domain.locations l on l.id = p.location_id
					WHERE p.id = $1`
	_psqlGetByUserId = `SELECT 
    				p.id,
    				l.id,
				    l.country,
				    l.city,
				    l.province,
				    l.district,
				    l.address,
				    l.lat,
				    l.long,
    				p.description,
    				p.type,
    				p.length,
    				p.width,
    				p.area,
    				p.floor,
    				p.number_of_floors,
    				p.rooms,
					p.bathrooms,
					p.yard,
					p.terrace,
					p.living_room,
					p.laundry_room,
					p.kitchen,
					p.garage,
    				u.id, 
    				u.user, 
    				u.password, 
    				u.email, 
    				u.theme,
    				u.person_id 
					FROM domain.properties p 
    				    INNER JOIN domain.users u ON p.user_id=u.id
    				    INNER JOIN domain.locations l on l.id = p.location_id
    				    WHERE p.user_id=$1`
	_psqlGetAll = `SELECT 
    				p.id,
    				l.id,
				    l.country,
				    l.city,
				    l.province,
				    l.district,
				    l.address,
				    l.lat,
				    l.long,
    				p.description,
    				p.type,
    				p.length,
    				p.width,
    				p.area,
    				p.floor,
    				p.number_of_floors,
    				p.rooms,
					p.bathrooms,
					p.yard,
					p.terrace,
					p.living_room,
					p.laundry_room,
					p.kitchen,
					p.garage,
    				u.id, 
    				u.user, 
    				u.password, 
    				u.email, 
    				u.theme,
    				u.person_id 
					FROM domain.properties p 
    				    INNER JOIN domain.users u ON p.user_id=u.id
    				    INNER JOIN domain.locations l on l.id = p.location_id`
	_psqlInsertProperty      = `INSERT INTO domain.properties (id, "user_id","location_id", "description", "type", "length","width","area","floor","number_of_floors","rooms","bathrooms","yard","terrace","living_room","laundry_room","kitchen","garage") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`
	_psqlInsertLocation      = `INSERT INTO domain.locations (id, country, city, province, district, address, lat, long) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_psqlInsertPropertyMedia = `INSERT INTO domain.properties_medias ("property_id","media_id") VALUES ($1, $2)`
	_psqlUpdateProperty      = `UPDATE domain.properties SET "user_id"=$2, "location_id"=$3, "description"=$4, "type"=$5, "length"=$6, "width"=$7, "area"=$8, "floor"=$9, "number_of_floors"=$10,"rooms"=$11,"bathrooms"=$12,"yard"=$13,"terrace"=$14,"living_room"=$15,"laundry_room"=$16,"kitchen"=$17, "garage"=$18 WHERE id=$1`
	_psqlUpdateLocation      = `UPDATE domain.locations SET "country"=$2, "city"=$3, "province"=$4, "district"=$5, "address"=$6, "lat"=$7, "long"=$8 WHERE id=$1`
	_psqlDelete              = `DELETE FROM domain.properties WHERE id=$1`
	_psqlDeleteMedias        = `DELETE FROM domain.properties_medias WHERE property_id=$1`
)

type Property struct {
	db *sql.DB
}

func New(db *sql.DB) Property {
	return Property{db}
}

func (p Property) GetStorageById(id uuid.UUID) (*model.PropertySecondLevel, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	m, err := p.scanRowWithUser(stmt.QueryRow(args...))
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (p Property) GetStorageAll() (model.PropertiesSecondLevel, error) {
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

	var ms model.PropertiesSecondLevel
	var m model.PropertySecondLevel

	for rows.Next() {
		m, err = p.scanRowWithUser(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Property) GetStorageByUserId(id uuid.UUID) (model.PropertiesSecondLevel, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlGetByUserId)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.PropertiesSecondLevel
	var m model.PropertySecondLevel

	for rows.Next() {
		m, err = p.scanRowWithUser(rows)
		if err != nil {
			break
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Property) CreateStorage(property model.Property) (*uuid.UUID, error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}
	property.ID = newId

	args := p.readModelProperty(property)

	stmt, err := p.db.Prepare(_psqlInsertProperty)
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

func (p Property) CreateCompleteStorage(property model.PropertyComplete, idsMedia []uuid.UUID) (*uuid.UUID, error) {
	idLocation, err := p.CreateLocation(property.Location)

	if err != nil {
		return nil, err
	}

	idProperty, err := p.CreateProperty(property.Property, idLocation)

	if err != nil {
		return nil, err
	}

	_, err = p.CreateMedias(idsMedia, idProperty)

	if err != nil {
		return nil, err
	}

	return idProperty, nil
}

func (p Property) CreateLocation(location model.Location) (*uuid.UUID, error) {
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

func (p Property) CreateProperty(property model.Property, idLocation *uuid.UUID) (*uuid.UUID, error) {
	newIdProperty, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Error generate UUID: %s\n", err)
	}

	property.ID = newIdProperty
	property.LocationID = *idLocation
	args := p.readModelProperty(property)

	stmt, err := p.db.Prepare(_psqlInsertProperty)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	return &newIdProperty, nil
}

func (p Property) CreateMedias(idsMedia []uuid.UUID, idProperty *uuid.UUID) (bool, error) {

	for _, element := range idsMedia {
		var propertyMedia model.PropertyMedia
		propertyMedia.PropertyID = *idProperty
		propertyMedia.MediaID = element

		args := p.readModelPropertyMedia(propertyMedia)

		stmt, err := p.db.Prepare(_psqlInsertPropertyMedia)
		if err != nil {
			return false, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(args...)
		if err != nil {
			return false, err
		}
	}
	return true, nil

}

func (p Property) UpdateStorage(id uuid.UUID, property model.Property) (bool, error) {
	property.ID = id

	args := p.readModelProperty(property)

	stmt, err := p.db.Prepare(_psqlUpdateProperty)
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

func (p Property) UpdateCompleteStorage(id uuid.UUID, propertyComplete model.PropertyComplete, idsMedia []uuid.UUID) (bool, error) {
	_, err := p.UpdateProperty(id, propertyComplete.Property)
	if err != nil {
		return false, err
	}

	_, err = p.UpdateLocation(propertyComplete.Property.LocationID, propertyComplete.Location)
	if err != nil {
		return false, err
	}

	_, err = p.UpdateMedia(id, idsMedia)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p Property) UpdateProperty(id uuid.UUID, property model.Property) (bool, error) {
	property.ID = id

	args := p.readModelProperty(property)
	stmt, err := p.db.Prepare(_psqlUpdateProperty)
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

func (p Property) UpdateLocation(id uuid.UUID, Location model.Location) (bool, error) {
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

func (p Property) UpdateMedia(id uuid.UUID, idsMedia []uuid.UUID) (bool, error) {
	_, err := p.DeleteMedia(id)
	if err != nil {
		return false, err
	}

	_, err = p.CreateMedias(idsMedia, &id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p Property) DeleteMedia(id uuid.UUID) (bool, error) {
	args := []any{id}

	stmt, err := p.db.Prepare(_psqlDeleteMedias)
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

func (p Property) DeleteStorage(id uuid.UUID) (bool, error) {
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

func (p Property) readModelProperty(property model.Property) []any {
	v := reflect.ValueOf(property)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (p Property) readModelLocation(location model.Location) []any {
	v := reflect.ValueOf(location)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (p Property) readModelPropertyMedia(propertyMedia model.PropertyMedia) []any {
	v := reflect.ValueOf(propertyMedia)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (p Property) scanRow(s pgx.Row) (model.Property, error) {
	m := model.Property{}

	err := s.Scan(
		&m.ID,
		&m.UserID,
		&m.Description,
		&m.Type,
		&m.Length,
		&m.Width,
		&m.Area,
		&m.Floor,
		&m.NumberOfFloors,
		&m.Rooms,
		&m.Bathrooms,
		&m.Yard,
		&m.Terrace,
		&m.LivingRoom,
		&m.LaundryRoom,
		&m.Kitchen,
		&m.Garage,
	)
	if err != nil {
		return model.Property{}, err
	}

	return m, nil
}

func (p Property) scanRowWithUser(s pgx.Row) (model.PropertySecondLevel, error) {
	m := model.PropertySecondLevel{}
	err := s.Scan(
		&m.ID,
		&m.Location.ID,
		&m.Location.Country,
		&m.Location.City,
		&m.Location.Province,
		&m.Location.District,
		&m.Location.Address,
		&m.Location.Lat,
		&m.Location.Long,
		&m.Description,
		&m.Type,
		&m.Length,
		&m.Width,
		&m.Area,
		&m.Floor,
		&m.NumberOfFloors,
		&m.Rooms,
		&m.Bathrooms,
		&m.Yard,
		&m.Terrace,
		&m.LivingRoom,
		&m.LaundryRoom,
		&m.Kitchen,
		&m.Garage,
		&m.User.ID,
		&m.User.User,
		&m.User.Password,
		&m.User.Email,
		&m.User.Theme,
		&m.User.PersonID,
	)
	if err != nil {
		return model.PropertySecondLevel{}, err
	}

	return m, nil
}
