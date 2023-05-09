package userrole

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseUserRole interface {
	GetByIDs(userID, roleID uuid.UUID) (*model.UserRoleOutput, error)
	GetAllByUserID(userID uuid.UUID) (model.UserRoleOutputs, error)
	Assignment(userID, roleID uuid.UUID) (*model.AssignOutput, error)
	Delete(userID, roleID uuid.UUID) (bool, error)
}
