package user

import (
	"fmt"
	"github.com/arthur8davis/housematch-api/application/repository/storage/user"
	"github.com/arthur8davis/housematch-api/domain/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage user.StorageUser
}

func New(storage user.StorageUser) User {
	return User{storage}
}

func (u User) GetById(id uuid.UUID) (*model.UserSecondLevel, error) {
	m, err := u.storage.GetByIdStorage(id)
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetById(): %w", err)
	}

	return m, nil
}

func (u User) GetAll() (model.UsersSecondLevel, error) {
	users, err := u.storage.GetAllStorage()
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetAll(): %w", err)
	}

	return users, nil
}

func (u User) GetAllWithRoles() (model.UsersWithRolesOutput, error) {
	users, err := u.storage.GetAllWithRolesStorage()
	if err != nil {
		return nil, fmt.Errorf("user.storage.GetAllWithRoles(): %w", err)
	}

	return users, nil
}

func (u User) Create(user model.User) (*model.CreateOutput, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("user.use.hashPassword(): %w", err)
	}
	user.Password = hashedPassword

	id, err := u.storage.CreateStorage(user)
	if err != nil {
		return nil, fmt.Errorf("user.use.Create(): %w", err)
		//return nil, fmt.Errorf("user.use.Create(): %s", err.Error())
	}

	var m model.CreateOutput
	m.Id = id

	return &m, nil
}

func (u User) Login(login model.Login) (*model.TokenOutput, error) {
	m, err := u.storage.GetByUsernameOrEmailStorage(login.User, login.Email)
	if err != nil {
		return nil, err
	}

	err = checkPassword(m.Password, login.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	tokenString, err := GenerateJWT(m.User, m.Email)
	if err != nil {
		return nil, err
	}

	var t model.TokenOutput
	t.Token = tokenString

	return &t, err
}

func (u User) Update(id uuid.UUID, user model.User) (*model.UpdateOutput, error) {
	created, err := u.storage.UpdateStorage(id, user)
	if err != nil {
		return nil, fmt.Errorf("user.storage.Update(): %w", err)
	}
	var m model.UpdateOutput
	m.Updated = created

	return &m, nil
}

func (u User) Delete(id uuid.UUID) (*model.DeleteOutput, error) {
	deleted, err := u.storage.DeleteStorage(id)
	if err != nil {
		return nil, fmt.Errorf("user.storage.Delete(): %w", err)
	}

	var m model.DeleteOutput
	m.Deleted = deleted

	return &m, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func checkPassword(password, passwordUser string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordUser)); err != nil {
		return err
	}

	return nil
}
