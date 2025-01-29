package services

import (
	"backend/internal/interfaces"
	"backend/internal/models"
	"github.com/google/uuid"
)

type UserService struct {
	Repo interfaces.UserRepositories
}

func NewUserService(repo interfaces.UserRepositories) interfaces.UserServices {
	return &UserService{Repo: repo}
}

func (u UserService) CreateUser(user *models.User) (*models.User, error) {
	return u.Repo.CreateUser(user)
}

func (u UserService) GetAllUsers() ([]*models.User, error) {
	return u.Repo.GetAllUsers()
}

func (u UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return u.Repo.GetUserByID(id)
}

func (u UserService) DeleteUser(id uuid.UUID) error {
	return u.Repo.DeleteUser(id)
}

func (u UserService) UpdateUser(id uuid.UUID, user *models.User) (*models.User, error) {
	return u.Repo.UpdateUser(id, user)
}
