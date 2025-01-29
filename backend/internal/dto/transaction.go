package dto

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type TransactionCreateRequest struct {
	Amount        float64    `json:"amount" validate:"required,gt=0"`
	Category      string     `json:"category" validate:"required"`
	Description   *string    `json:"description"`
	BankAccountID *uuid.UUID `json:"bank_account_id"`
	CreditCardID  *uuid.UUID `json:"credit_card_id"`
}

type TransactionUpdateRequest struct {
	ID            uuid.UUID  `json:"id" validate:"required"`
	UserID        uuid.UUID  `json:"user_id" validate:"required"`
	BankAccountID *uuid.UUID `json:"bank_account_id"`
	CreditCardID  *uuid.UUID `json:"credit_card_id"`
	Amount        float64    `json:"amount" validate:"required,gt=0"`
	Category      string     `json:"category" validate:"required"`
	Description   *string    `json:"description"`
}
type TransactionResponseList struct {
	Data    []*models.Transaction `json:"data"`
	Code    string                `json:"code"`
	Message string                `json:"message"`
}

type TransactionResponse struct {
	Data    *models.Transaction `json:"data"`
	Code    string              `json:"code"`
	Message string              `json:"message"`
}
