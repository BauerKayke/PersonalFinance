package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type BankAccountRepositories interface {
	CreateBankAccount(bankAccount *models.BankAccount) (*models.BankAccount, error)
	GetAllBankAccounts(userId *uuid.UUID) ([]*models.BankAccount, error)
	GetBankAccountByID(bankAccountId, userId *uuid.UUID) (*models.BankAccount, error)
	DeleteBankAccount(bankAccountId uuid.UUID) error
	UpdateBankAccount(bankAccountId uuid.UUID, bankAccount *models.BankAccount) (*models.BankAccount, error)
}
