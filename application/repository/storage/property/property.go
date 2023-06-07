package property

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageProperty interface {
	GetByIdStorage(id uuid.UUID) (*model.PropertySecondLevel, error)
	GetAllStorage() (model.PropertiesSecondLevel, error)
	GetByUserIdStorage(id uuid.UUID) (model.PropertiesSecondLevel, error)
	CreateStorage(m model.Property) (*uuid.UUID, error)
	CreateCompleteStorage(m model.PropertyComplete, idsMedia []uuid.UUID) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, model model.Property) (bool, error)
	UpdateCompleteStorage(id uuid.UUID, model model.PropertyComplete, idsMedia []uuid.UUID) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
