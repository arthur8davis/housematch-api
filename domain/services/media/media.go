package media

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseMedia interface {
	Create(m model.Media) (*model.CreateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
