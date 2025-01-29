package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type BankAccountServices interface {
	CreateBankAccount(bankAccount *models.BankAccount) (*models.BankAccount, error)
	GetAllBankAccounts() ([]*models.BankAccount, error)
	GetBankAccountByID(id uuid.UUID) (*models.BankAccount, error)
	DeleteBankAccount(id uuid.UUID) error
	UpdateBankAccount(id uuid.UUID, bankAccount *models.BankAccount) (*models.BankAccount, error)
}
