package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
	"github.com/google/uuid"
	"time"
)

func ToUserModel(dto dto.UserCreateRequest) models.User {
	return models.User{
		ID:        uuid.New(),
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToUserUpdateModel(dto dto.UserUpdatedRequest) models.User {
	return models.User{
		ID:        dto.ID,
		Email:     dto.Email,
		Password:  dto.Password,
		UpdatedAt: time.Now(),
	}
}
