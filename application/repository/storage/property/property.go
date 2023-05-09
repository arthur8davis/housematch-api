package property

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageProperty interface {
	GetStorageById(id uuid.UUID) (*model.PropertySecondLevel, error)
	GetStorageAll() (model.PropertiesSecondLevel, error)
	GetStorageByUserId(id uuid.UUID) (model.PropertiesSecondLevel, error)
	CreateStorage(m model.Property) (*uuid.UUID, error)
	CreateCompleteStorage(m model.PropertyComplete, idsMedia []uuid.UUID) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, model model.Property) (bool, error)
	UpdateCompleteStorage(id uuid.UUID, model model.PropertyComplete, idsMedia []uuid.UUID) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
