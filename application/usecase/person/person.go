package person

import (
	"fmt"
	"github.com/Melany751/house-match-server/application/repository/storage/person"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type Person struct {
	storage person.StoragePerson
}

func New(storage person.StoragePerson) Person {
	return Person{storage}
}

func (p Person) GetById(id uuid.UUID) (*model.PersonSecondLevel, error) {
	m, err := p.storage.GetByIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("person.storage.GetById(): %w", err)
	}

	return m, nil
}

func (p Person) GetAll() (model.PersonsOutput, error) {
	ms, err := p.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("person.storage.GetAll(): %w", err)
	}

	return ms, nil
}

func (p Person) Create(person model.Person) (*model.CreateOutput, error) {
	id, err := p.storage.CreateStorage(person)
	if err != nil {
		return nil, fmt.Errorf("person.use.Create(): %w", err)
		//return nil, fmt.Errorf("person.use.Create(): %s", err.Error())
	}

	var m model.CreateOutput
	m.Id = id

	return &m, nil
}

func (p Person) Update(id uuid.UUID, person model.Person) (*model.UpdateOutput, error) {
	created, err := p.storage.UpdateStorage(id, person)
	if err != nil {
		return nil, fmt.Errorf("person.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = created

	return &m, nil
}

func (p Person) Delete(id uuid.UUID) (*model.DeleteOutput, error) {
	deleted, err := p.storage.DeleteStorage(id)
	if err != nil {
		return nil, fmt.Errorf("person.storage.Delete(): %w", err)
	}

	var m model.DeleteOutput
	m.Deleted = deleted

	return &m, nil
}
