package transaction

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseTransaction interface {
	GetById(id uuid.UUID) (*model.TransactionSecondLevel, error)
	GetByUserId(id uuid.UUID) (model.TransactionsSecondLevel, error)
	GetAll() (model.TransactionsSecondLevel, error)
	GetAllByFilters(params map[string]string) (model.TransactionsThirdLevel, error)
	Create(m model.Transaction) (*model.CreateOutput, error)
	Update(id uuid.UUID, model model.Transaction) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (bool, error)
}
