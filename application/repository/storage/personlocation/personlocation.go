package personlocation

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StoragePersonLocation interface {
	CreateStorage(m model.PersonLocation) (*uuid.UUID, error)
	UpdateStorage(personID uuid.UUID, user model.PersonLocation) (bool, error)
}
