package useruserRole

import (
	"fmt"
	userRole "github.com/arthur8davis/housematch-api/application/repository/storage/userrole"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UserRole struct {
	storage userRole.StorageUserRole
}

func New(storage userRole.StorageUserRole) UserRole {
	return UserRole{storage}
}

func (r UserRole) GetByIDs(userID, roleID uuid.UUID) (*model.UserRoleOutput, error) {
	fmt.Println("#######################a")
	user, err := r.storage.GetByIDsStorage(userID, roleID)
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetById(): %w", err)
	}
	fmt.Println("#######################b")

	return user, nil
}

func (r UserRole) GetAllByUserID(userID uuid.UUID) (model.UserRoleOutputs, error) {
	userRoles, err := r.storage.GetAllByUserIDStorage(userID)
	if err != nil {
		return nil, fmt.Errorf("userRole.storage.GetAll(): %w", err)
	}
	return userRoles, nil
}

func (r UserRole) Assignment(userID, roleID uuid.UUID) (*model.AssignOutput, error) {
	id, err := r.storage.AssignmentStorage(userID, roleID)
	if err != nil {
		return nil, fmt.Errorf("userRole.storage.Create(): %w", err)
	}
	var m = model.AssignOutput{Assigned: id}
	return &m, nil
}

func (r UserRole) Delete(userID, roleID uuid.UUID) (bool, error) {
	deleted, err := r.storage.DeleteStorage(userID, roleID)
	if err != nil {
		return false, fmt.Errorf("userRole.storage.Delete(): %w", err)
	}

	return deleted, nil
}
