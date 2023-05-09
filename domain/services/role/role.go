package role

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseRole interface {
	GetById(id uuid.UUID) (*model.Role, error)
	GetAll() (model.Roles, error)
	Create(m model.Role) (*model.CreateOutput, error)
	Update(id uuid.UUID, model model.Role) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (bool, error)
}
