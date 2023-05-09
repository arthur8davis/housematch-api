package roleview

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseRoleView interface {
	GetByIDs(roleID, viewID uuid.UUID) (*model.RoleViewOutput, error)
	GetAll() (model.RoleViewOutputs, error)
	Assignment(roleView model.RoleView) (bool, error)
	Update(roleView model.RoleView) (*model.UpdateOutput, error)
	Delete(roleID, viewID uuid.UUID) (bool, error)
}
