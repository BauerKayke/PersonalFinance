package models

import (
	"github.com/google/uuid"
	"time"
)

type BankAccount struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	BankName  string    `json:"bank_name"`
	AccountNo string    `json:"account_no"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
