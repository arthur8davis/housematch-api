package media

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageMedia interface {
	CreateStorage(m model.Media) error
	DeleteStorage(id uuid.UUID) (bool, error)
}
