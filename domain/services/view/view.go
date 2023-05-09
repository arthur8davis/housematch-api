package view

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseView interface {
	GetById(id uuid.UUID) (*model.ViewOutput, error)
	GetAll() (model.ViewsOutput, error)
	Create(m model.View) (*model.CreateOutput, error)
	Update(id uuid.UUID, model model.View) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (bool, error)
}
