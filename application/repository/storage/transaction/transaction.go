package transaction

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type StorageTransaction interface {
	GetStorageById(id uuid.UUID) (*model.TransactionSecondLevel, error)
	GetStorageByUserId(id uuid.UUID) (model.TransactionsSecondLevel, error)
	GetStorageAll() (model.TransactionsSecondLevel, error)
	GetStorageAllByFilters(params map[string]string) (model.TransactionsThirdLevel, error)
	CreateStorage(m model.Transaction) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, role model.Transaction) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
