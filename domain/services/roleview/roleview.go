package roleview

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseRoleView interface {
	GetByIDs(roleID, viewID uuid.UUID) (*model.RoleViewOutput, error)
	GetAll() (model.RoleViewOutputs, error)
	Assignment(roleView model.RoleView) (bool, error)
	Update(roleView model.RoleView) (*model.UpdateOutput, error)
	Delete(roleID, viewID uuid.UUID) (bool, error)
}
