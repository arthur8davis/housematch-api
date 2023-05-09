package person

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StoragePerson interface {
	GetByIdStorage(id uuid.UUID) (*model.PersonSecondLevel, error)
	GetAllStorage() (model.PersonsOutput, error)
	CreateStorage(m model.Person) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, user model.Person) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
