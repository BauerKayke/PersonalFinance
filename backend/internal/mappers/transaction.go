package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
	"github.com/google/uuid"
	"time"
)

// Converter TransactionCreateRequest para Transaction (Model)
func ToTransactionModel(dto dto.TransactionCreateRequest, userID *uuid.UUID) models.Transaction {
	return models.Transaction{
		ID:            uuid.New(),
		UserID:        *userID,
		BankAccountID: dto.BankAccountID,
		CreditCardID:  dto.CreditCardID,
		Amount:        dto.Amount,
		Category:      dto.Category,
		Description:   dto.Description,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

// Converter TransactionUpdateRequest para Transaction (Model)
func ToTransactionUpdateModel(dto dto.TransactionUpdateRequest, userID *uuid.UUID) models.Transaction {
	return models.Transaction{
		ID:            dto.ID,
		UserID:        *userID,
		BankAccountID: dto.BankAccountID,
		CreditCardID:  dto.CreditCardID,
		Amount:        dto.Amount,
		Category:      dto.Category,
		Description:   dto.Description,
		UpdatedAt:     time.Now(),
	}
}
