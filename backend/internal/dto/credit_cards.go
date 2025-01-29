package dto

import (
	"backend/internal/models"
	"github.com/google/uuid"
	"time"
)

type CreditCardCreateRequest struct {
	UserID        uuid.UUID  `json:"user_id"`
	BankAccountID *uuid.UUID `json:"bank_account_id" validate:"required"`
	CardName      string     `json:"card_name" validate:"required"`
	CardNumber    string     `json:"card_number" validate:"required"`
	CreditLimit   float64    `json:"credit_limit" validate:"required"`
	Available     float64    `json:"available" validate:"required"`
	Expiration    time.Time  `json:"expiration" validate:"required"`
}

type CreditCardUpdateRequest struct {
	ID            uuid.UUID  `json:"id" validate:"required"`
	UserID        uuid.UUID  `json:"user_id" validate:"required"`
	BankAccountID *uuid.UUID `json:"bank_account_id"`
	CardName      string     `json:"card_name" validate:"required"`
	CardNumber    string     `json:"card_number" validate:"required"`
	CreditLimit   float64    `json:"credit_limit" validate:"required"`
	Available     float64    `json:"available" validate:"required"`
	Expiration    time.Time  `json:"expiration" validate:"required"`
}

type CreditCardResponse struct {
	Data    *models.CreditCard `json:"data"`
	Code    string             `json:"code"`
	Message string             `json:"message"`
}

type CreditCardResponseList struct {
	Data    []*models.CreditCard `json:"data"`
	Code    string               `json:"code"`
	Message string               `json:"message"`
}
