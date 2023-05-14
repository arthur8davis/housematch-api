package location

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/application/repository/storage/locationperson"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type Location struct {
	storage locationperson.StorageLocation
}

func New(storage locationperson.StorageLocation) Location {
	return Location{storage}
}

func (l Location) GetById(id uuid.UUID) (*model.Location, error) {
	user, err := l.storage.GetByIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.GetById(): %w", err)
	}

	return user, nil
}

func (l Location) GetAll() (model.Locations, error) {
	users, err := l.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.GetAll(): %w", err)
	}

	return users, nil
}

func (l Location) Create(location model.Location) (*model.CreateOutput, error) {
	id, err := l.storage.CreateStorage(location)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.use.Create(): %w", err)
		//return nil, fmt.Errorf("locationPerson.use.Create(): %s", err.Error())
	}

	var m model.CreateOutput
	m.Id = id

	return &m, nil
}

func (l Location) Update(id uuid.UUID, location model.Location) (*model.UpdateOutput, error) {
	created, err := l.storage.UpdateStorage(id, location)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = created

	return &m, nil
}

func (l Location) Delete(id uuid.UUID) (*model.DeleteOutput, error) {
	deleted, err := l.storage.DeleteStorage(id)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.Delete(): %w", err)
	}

	var m model.DeleteOutput
	m.Deleted = deleted

	return &m, nil
}
