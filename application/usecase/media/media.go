package media

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/application/repository/storage/media"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type Media struct {
	storage media.StorageMedia
}

func New(storage media.StorageMedia) Media {
	return Media{storage}
}

func (m Media) Create(file model.Media) (*model.CreateOutput, error) {
	err := m.storage.CreateStorage(file)
	if err != nil {
		return nil, fmt.Errorf("media.use.Create(): %w", err)
	}

	return &model.CreateOutput{Id: &file.ID}, nil
}

func (m Media) Delete(id uuid.UUID) (*model.DeleteOutput, error) {
	deleted, err := m.storage.DeleteStorage(id)
	if err != nil {
		return nil, fmt.Errorf("media.storage.Delete(): %w", err)
	}

	return &model.DeleteOutput{Deleted: deleted}, err
}
