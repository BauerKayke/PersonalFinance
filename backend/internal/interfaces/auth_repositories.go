package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type AuthRepositories interface {
	GetUserByEmail(email string) (*models.User, error)
	SaveSession(session *models.Sessions) (bool, error)
	GetActiveSession(userId *uuid.UUID) (bool, error)
	DeleteSession(token string) error
}
