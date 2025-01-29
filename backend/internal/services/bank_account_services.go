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

func (b BankAccountService) GetAllBankAccounts() ([]*models.BankAccount, error) {
	return b.Repo.GetAllBankAccounts()
}

func (b BankAccountService) GetBankAccountByID(id uuid.UUID) (*models.BankAccount, error) {
	return b.Repo.GetBankAccountByID(id)
}

func (b BankAccountService) DeleteBankAccount(id uuid.UUID) error {
	return b.Repo.DeleteBankAccount(id)
}

func (b BankAccountService) UpdateBankAccount(id uuid.UUID, bankAccount *models.BankAccount) (*models.BankAccount, error) {
	return b.Repo.UpdateBankAccount(id, bankAccount)
}
