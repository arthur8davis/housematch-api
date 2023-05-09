package module

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageModule interface {
	GetStorageById(id uuid.UUID) (*model.Module, error)
	GetStorageAll() (model.Modules, error)
	CreateStorage(m model.Module) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, module model.Module) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
