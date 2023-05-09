package locationperson

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageLocation interface {
	GetByIdStorage(id uuid.UUID) (*model.Location, error)
	GetAllStorage() (model.Locations, error)
	CreateStorage(m model.Location) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, user model.Location) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
