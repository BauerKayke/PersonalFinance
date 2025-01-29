package dto

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type BankAccountCreateRequest struct {
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	BankName  string    `json:"bank_name" validate:"required"`
	AccountNo string    `json:"account_no" validate:"required"`
	Balance   float64   `json:"balance" validate:"required"`
}

type BankAccountUpdateRequest struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	BankName  string    `json:"bank_name" validate:"required"`
	AccountNo string    `json:"account_no" validate:"required"`
	Balance   float64   `json:"balance" validate:"required"`
}

type BankAccountResponse struct {
	Data    *models.BankAccount `json:"data"`
	Code    string              `json:"code"`
	Message string              `json:"message"`
}
type BankAccountResponseList struct {
	Data    []*models.BankAccount `json:"data"`
	Code    string                `json:"code"`
	Message string                `json:"message"`
}
