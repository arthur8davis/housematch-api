package view

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageView interface {
	GetStorageById(id uuid.UUID) (*model.ViewOutput, error)
	GetStorageAll() (model.ViewsOutput, error)
	CreateStorage(m model.View) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, view model.View) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
