package role

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageRole interface {
	GetStorageById(id uuid.UUID) (*model.Role, error)
	GetStorageAll() (model.Roles, error)
	CreateStorage(m model.Role) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, role model.Role) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
