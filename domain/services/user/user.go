package user

import (
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
)

type UseCaseUser interface {
	Login(login model.Login) (*model.TokenOutput, error)
	GetById(id uuid.UUID) (*model.UserSecondLevel, error)
	GetAll() (model.UsersSecondLevel, error)
	GetAllWithRoles() (model.UsersWithRolesOutput, error)
	Create(m model.User) (*model.CreateOutput, error)
	Update(id uuid.UUID, user model.User) (*model.UpdateOutput, error)
	Delete(id uuid.UUID) (*model.DeleteOutput, error)
}
