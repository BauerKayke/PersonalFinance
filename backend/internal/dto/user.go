package dto

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdatedRequest struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	Data    *models.User `json:"data"`
	Code    string       `json:"code"`
	Message string       `json:"message"`
}
type UserResponseList struct {
	Data    []*models.User `json:"data"`
	Code    string         `json:"code"`
	Message string         `json:"message"`
}
