package user

import (
	"github.com/Melany751/house-match-server/domain/model"
	"github.com/google/uuid"
)

type StorageUser interface {
	GetByIdStorage(id uuid.UUID) (*model.UserSecondLevel, error)
	GetByUsernameOrEmailStorage(username, email string) (*model.User, error)
	GetAllStorage() (model.UsersSecondLevel, error)
	GetAllWithRolesStorage() (model.UsersWithRolesOutput, error)
	CreateStorage(m model.User) (*uuid.UUID, error)
	UpdateStorage(id uuid.UUID, user model.User) (bool, error)
	DeleteStorage(id uuid.UUID) (bool, error)
}
