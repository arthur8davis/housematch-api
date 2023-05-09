package userrole

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageUserRole interface {
	GetByIDsStorage(userID, roleID uuid.UUID) (*model.UserRoleOutput, error)
	GetAllByUserIDStorage(userID uuid.UUID) (model.UserRoleOutputs, error)
	AssignmentStorage(userID, roleID uuid.UUID) (bool, error)
	DeleteStorage(userID, roleID uuid.UUID) (bool, error)
}
