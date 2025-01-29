package models

import (
	"github.com/google/uuid"
	"time"
)

type CreditCard struct {
	ID            uuid.UUID  `json:"id"`
	UserID        uuid.UUID  `json:"user_id"`
	BankAccountID *uuid.UUID `json:"bank_account_id"`
	CardName      string     `json:"card_name"`
	CardNumber    string     `json:"card_number"`
	CreditLimit   float64    `json:"credit_limit"`
	Available     float64    `json:"available"`
	Expiration    time.Time  `json:"expiration"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
