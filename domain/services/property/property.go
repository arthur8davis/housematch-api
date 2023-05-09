package property

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseModule interface {
	GetById(id uuid.UUID) (*model.PropertySecondLevel, error)
	GetAll() (model.PropertiesSecondLevel, error)
	GetByUserId(id uuid.UUID) (model.PropertiesSecondLevel, error)
	Create(m model.Property) (*model.CreateOutput, error)
	CreateComplete(m model.PropertyComplete, idsMedia []uuid.UUID) (*model.CreateOutput, error)
	Update(id uuid.UUID, model model.Property) (*model.UpdateOutput, error)
	UpdateComplete(id uuid.UUID, model model.PropertyComplete, idsMedia []uuid.UUID) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (bool, error)
}
