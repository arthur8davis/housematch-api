package personlocation

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StoragePersonLocation interface {
	CreateStorage(m model.PersonLocation) (*uuid.UUID, error)
	UpdateStorage(personID uuid.UUID, user model.PersonLocation) (bool, error)
}
