package role

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageRole interface {
	GetStorageById(id uuid.UUID) (*model.Role, error)
	GetStorageAll() (model.Roles, error)
	CreateStorage(m model.Role) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, role model.Role) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
