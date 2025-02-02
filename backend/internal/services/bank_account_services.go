package services

import (
	"backend/internal/interfaces"
	"backend/internal/models"
	"github.com/google/uuid"
)

type BankAccountService struct {
	Repo interfaces.BankAccountRepositories
}

func NewBankAccountService(repo interfaces.BankAccountRepositories) interfaces.BankAccountServices {
	return &BankAccountService{Repo: repo}
}

func (b BankAccountService) CreateBankAccount(bankAccount *models.BankAccount) (*models.BankAccount, error) {
	return b.Repo.CreateBankAccount(bankAccount)
}

func (b BankAccountService) GetAllBankAccounts(userId *uuid.UUID) ([]*models.BankAccount, error) {
	return b.Repo.GetAllBankAccounts(userId)
}

func (b BankAccountService) GetBankAccountByID(bankAccountId, userId *uuid.UUID) (*models.BankAccount, error) {
	return b.Repo.GetBankAccountByID(bankAccountId, userId)
}

func (b BankAccountService) DeleteBankAccount(bankAccountId uuid.UUID) error {
	return b.Repo.DeleteBankAccount(bankAccountId)
}

func (b BankAccountService) UpdateBankAccount(bankAccountId uuid.UUID, bankAccount *models.BankAccount) (*models.BankAccount, error) {
	return b.Repo.UpdateBankAccount(bankAccountId, bankAccount)
}
