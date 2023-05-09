package person

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCasePerson interface {
	GetById(id uuid.UUID) (*model.PersonSecondLevel, error)
	GetAll() (model.PersonsOutput, error)
	Create(m model.Person) (*model.CreateOutput, error)
	Update(id uuid.UUID, user model.Person) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
