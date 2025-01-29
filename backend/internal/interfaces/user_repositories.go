package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type UserRepositories interface {
	CreateUser(user *models.User) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user *models.User) (*models.User, error)
}
