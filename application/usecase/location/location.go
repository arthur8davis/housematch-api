package location

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/locationperson"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Location struct {
	storage locationperson.StorageLocation
}

func New(storage locationperson.StorageLocation) Location {
	return Location{storage}
}

func (u Location) GetById(id uuid.UUID) (*model.Location, error) {
	user, err := u.storage.GetByIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.GetById(): %w", err)
	}

	return user, nil
}

func (u Location) GetAll() (model.Locations, error) {
	users, err := u.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.GetAll(): %w", err)
	}

	return users, nil
}

func (u Location) Create(user model.Location) (*model.CreateOutput, error) {
	id, err := u.storage.CreateStorage(user)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.use.Create(): %w", err)
		//return nil, fmt.Errorf("locationPerson.use.Create(): %s", err.Error())
	}

	var m model.CreateOutput
	m.Id = id

	return &m, nil
}

func (u Location) Update(id uuid.UUID, user model.Location) (*model.UpdateOutput, error) {
	created, err := u.storage.UpdateStorage(id, user)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = created

	return &m, nil
}

func (u Location) Delete(id uuid.UUID) (*model.DeleteOutput, error) {
	deleted, err := u.storage.DeleteStorage(id)
	if err != nil {
		return nil, fmt.Errorf("locationPerson.storage.Delete(): %w", err)
	}

	var m model.DeleteOutput
	m.Deleted = deleted

	return &m, nil
}
