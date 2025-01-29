package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
	"github.com/google/uuid"
	"time"
)

// Converter BankAccountCreateRequest para BankAccount (Model)
func ToBankAccountModel(dto dto.BankAccountCreateRequest, userID *uuid.UUID) models.BankAccount {
	return models.BankAccount{
		ID:        uuid.New(),
		UserID:    *userID,
		BankName:  dto.BankName,
		AccountNo: dto.AccountNo,
		Balance:   dto.Balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Converter BankAccountUpdateRequest para BankAccount (Model)
func ToBankAccountUpdateModel(dto dto.BankAccountUpdateRequest, userID *uuid.UUID) models.BankAccount {
	return models.BankAccount{
		ID:        dto.ID,
		UserID:    *userID,
		BankName:  dto.BankName,
		AccountNo: dto.AccountNo,
		Balance:   dto.Balance,
		UpdatedAt: time.Now(),
	}
}
