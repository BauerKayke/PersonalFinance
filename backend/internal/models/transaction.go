package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            uuid.UUID  `json:"id"`
	UserID        uuid.UUID  `json:"user_id"`
	BankAccountID *uuid.UUID `json:"bank_account_id"`
	CreditCardID  *uuid.UUID `json:"credit_card_id"`
	Amount        float64    `json:"amount"`
	Category      string     `json:"category"`
	Description   *string    `json:"description"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
