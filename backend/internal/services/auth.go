package services

import (
	"backend/internal/interfaces"
	"backend/internal/models"
	"github.com/google/uuid"
)

type AuthService struct {
	Repo interfaces.AuthRepositories
}

func NewAuthService(repo interfaces.AuthRepositories) interfaces.AuthServices {
	return &AuthService{Repo: repo}
}

func (u AuthService) GetUserByEmail(username string) (*models.User, error) {
	return u.Repo.GetUserByEmail(username)
}

func (u AuthService) SaveSession(session *models.Sessions) (bool, error) {
	return u.Repo.SaveSession(session)
}

func (u AuthService) DeleteSession(userId *uuid.UUID) error {
	return u.Repo.DeleteSession(userId)
}

func (u AuthService) GetActiveSession(userId *uuid.UUID) (bool, error) {
	return u.Repo.GetActiveSession(userId)
}
