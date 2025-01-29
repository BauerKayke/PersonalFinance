package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type TransactionRepositories interface {
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	GetAllTransactions() ([]*models.Transaction, error)
	GetTransactionByID(id uuid.UUID) (*models.Transaction, error)
	DeleteTransaction(id uuid.UUID) error
	UpdateTransaction(id uuid.UUID, transaction *models.Transaction) (*models.Transaction, error)
}
