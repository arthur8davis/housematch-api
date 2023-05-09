package module

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/application/repository/storage/module"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type Module struct {
	storage module.StorageModule
}

func New(storage module.StorageModule) Module {
	return Module{storage}
}

func (m Module) GetById(id uuid.UUID) (*model.Module, error) {
	module, err := m.storage.GetStorageById(id)
	if err != nil {
		return nil, fmt.Errorf("module.storage.GetById(): %w", err)
	}

	return module, nil
}

func (m Module) GetAll() (model.Modules, error) {
	modules, err := m.storage.GetStorageAll()
	if err != nil {
		return nil, fmt.Errorf("module.storage.GetAll(): %w", err)
	}

	return modules, nil
}

func (m Module) Create(module model.Module) (*model.CreateOutput, error) {
	id, err := m.storage.CreateStorage(module)
	if err != nil {
		return nil, fmt.Errorf("module.storage.Create(): %w", err)
	}

	var mo model.CreateOutput
	mo.Id = id

	return &mo, nil
}

func (m Module) Update(id uuid.UUID, module model.Module) (*model.UpdateOutput, error) {
	updated, err := m.storage.UpdateStorage(id, module)
	if err != nil {
		return nil, fmt.Errorf("module.storage.Update(): %w", err)
	}

	var mo model.UpdateOutput
	mo.Updated = updated

	return &mo, nil
}

func (m Module) Delete(id uuid.UUID) (bool, error) {
	deleted, err := m.storage.DeleteStorage(id)
	if err != nil {
		return false, fmt.Errorf("module.storage.Delete(): %w", err)
	}

	return deleted, nil
}
