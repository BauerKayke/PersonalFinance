package services

import (
	"backend/internal/interfaces"
	"backend/internal/models"
	"github.com/google/uuid"
)

type TransactionService struct {
	Repo interfaces.TransactionRepositories
}

func NewTransactionService(repo interfaces.TransactionRepositories) interfaces.TransactionServices {
	return &TransactionService{Repo: repo}
}

func (t TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return t.Repo.CreateTransaction(transaction)
}

func (t TransactionService) GetAllTransactions(userID *uuid.UUID) ([]*models.Transaction, error) {
	return t.Repo.GetAllTransactions(userID)
}

func (t TransactionService) GetTransactionByID(transactionID, userID *uuid.UUID) (*models.Transaction, error) {
	return t.Repo.GetTransactionByID(transactionID, userID)
}

func (t TransactionService) DeleteTransaction(id uuid.UUID) error {
	return t.Repo.DeleteTransaction(id)
}

func (t TransactionService) UpdateTransaction(id uuid.UUID, transaction *models.Transaction) (*models.Transaction, error) {
	return t.Repo.UpdateTransaction(id, transaction)
}
