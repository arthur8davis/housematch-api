package media

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type UseCaseMedia interface {
	GetById(id uuid.UUID) (*model.Media, error)
	GetAll() (model.Medias, error)
	Create(m model.Media) (*model.CreateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
