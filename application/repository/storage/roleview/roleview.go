package roleview

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageRoleView interface {
	AssignmentStorage(roleView model.RoleView) (bool, error)
	GetByIDsStorage(roleID, viewID uuid.UUID) (*model.RoleViewOutput, error)
	GetAllStorage() (model.RoleViewOutputs, error)
	UpdateStorage(roleView model.RoleView) (bool, error)
	DeleteStorage(roleID, viewID uuid.UUID) (bool, error)
}
