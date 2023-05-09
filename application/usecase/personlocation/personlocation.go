package personlocation

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/application/repository/storage/personlocation"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type PersonLocation struct {
	storage personlocation.StoragePersonLocation
}

func New(storage personlocation.StoragePersonLocation) PersonLocation {
	return PersonLocation{storage}
}

func (p PersonLocation) Create(personLocation model.PersonLocation) (*model.CreateOutput, error) {
	id, err := p.storage.CreateStorage(personLocation)
	if err != nil {
		return nil, fmt.Errorf("personLocation.use.Create(): %w", err)
		//return nil, fmt.Errorf("person.use.Create(): %s", err.Error())
	}

	var m model.CreateOutput
	m.Id = id

	return &m, nil
}

func (p PersonLocation) Update(personID uuid.UUID, personLocation model.PersonLocation) (*model.UpdateOutput, error) {
	created, err := p.storage.UpdateStorage(personID, personLocation)
	if err != nil {
		return nil, fmt.Errorf("personLocation.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = created

	return &m, nil
}
