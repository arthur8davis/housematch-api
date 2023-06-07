package property

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/application/repository/storage/property"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type Property struct {
	storage property.StorageProperty
}

func New(storage property.StorageProperty) Property {
	return Property{storage}
}

func (p Property) GetById(id uuid.UUID) (*model.PropertySecondLevel, error) {
	property, err := p.storage.GetByIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("property.storage.GetById(): %w", err)
	}

	return property, nil
}

func (p Property) GetAll() (model.PropertiesSecondLevel, error) {
	properties, err := p.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("property.storage.GetAll(): %w", err)
	}

	return properties, nil
}

func (p Property) GetByUserId(id uuid.UUID) (model.PropertiesSecondLevel, error) {
	properties, err := p.storage.GetByUserIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("property.storage.GetAll(): %w", err)
	}

	return properties, nil
}

func (p Property) Create(property model.Property) (*model.CreateOutput, error) {
	id, err := p.storage.CreateStorage(property)
	if err != nil {
		return nil, fmt.Errorf("property.storage.Create(): %w", err)
	}
	var m model.CreateOutput
	m.Id = id
	return &m, nil
}

func (p Property) CreateComplete(propertyComplete model.PropertyComplete, idsMedia []uuid.UUID) (*model.CreateOutput, error) {
	id, err := p.storage.CreateCompleteStorage(propertyComplete, idsMedia)
	if err != nil {
		return nil, fmt.Errorf("property.storage.CreateComplete(): %w", err)
	}
	var m model.CreateOutput
	m.Id = id
	return &m, nil
}

func (p Property) Update(id uuid.UUID, property model.Property) (*model.UpdateOutput, error) {
	updated, err := p.storage.UpdateStorage(id, property)
	if err != nil {
		return nil, fmt.Errorf("property.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = updated
	return &m, nil
}

func (p Property) UpdateComplete(id uuid.UUID, propertyComplete model.PropertyComplete, idsMedia []uuid.UUID) (*model.UpdateOutput, error) {
	updated, err := p.storage.UpdateCompleteStorage(id, propertyComplete, idsMedia)
	if err != nil {
		return nil, fmt.Errorf("property.storage.UpdateComplete(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = updated
	return &m, nil
}

func (p Property) Delete(id uuid.UUID) (bool, error) {
	deleted, err := p.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("property.storage.Delete(): %w", err)
	}

	return deleted, nil
}
