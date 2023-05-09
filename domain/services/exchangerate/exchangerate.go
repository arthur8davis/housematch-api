package exchangerate

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseExchangeRate interface {
	GetById(id uuid.UUID) (*model.ExchangeRate, error)
	GetAll() (model.ExchangeRate, error)
	Create(m model.ExchangeRate) (*model.CreateOutput, error)
	Update(id uuid.UUID, user model.ExchangeRate) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
