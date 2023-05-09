package person

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StoragePerson interface {
	GetByIdStorage(id uuid.UUID) (*model.PersonSecondLevel, error)
	GetAllStorage() (model.PersonsOutput, error)
	CreateStorage(m model.Person) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, user model.Person) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
