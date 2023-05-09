package roleview

import (
	"fmt"
	roleView "github.com/Melany751/house-match-server/application/repository/storage/roleview"
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type RoleView struct {
	storage roleView.StorageRoleView
}

func New(storage roleView.StorageRoleView) RoleView {
	return RoleView{storage}
}

func (rv RoleView) GetByIDs(roleID, viewID uuid.UUID) (*model.RoleViewOutput, error) {
	roleView, err := rv.storage.GetByIDsStorage(roleID, viewID)
	if err != nil {
		return nil, fmt.Errorf("roleView.storage.GetById(): %w", err)
	}

	return roleView, nil
}

func (rv RoleView) GetAll() (model.RoleViewOutputs, error) {
	roleViews, err := rv.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("roleView.storage.GetAll(): %w", err)
	}

	return roleViews, nil
}

func (rv RoleView) Assignment(roleView model.RoleView) (bool, error) {
	assignment, err := rv.storage.AssignmentStorage(roleView)
	if err != nil {
		return false, fmt.Errorf("roleView.storage.Create(): %w", err)
	}

	return assignment, nil
}

func (rv RoleView) Update(roleView model.RoleView) (*model.UpdateOutput, error) {
	updated, err := rv.storage.UpdateStorage(roleView)
	if err != nil {
		return nil, fmt.Errorf("roleView.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = updated
	return &m, nil
}

func (rv RoleView) Delete(roleID, viewID uuid.UUID) (bool, error) {
	deleted, err := rv.storage.DeleteStorage(roleID, viewID)
	if err != nil {
		return false, fmt.Errorf("roleView.storage.Delete(): %w", err)
	}

	return deleted, nil
}
