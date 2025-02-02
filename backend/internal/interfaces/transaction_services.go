package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type TransactionServices interface {
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	GetAllTransactions(userID *uuid.UUID) ([]*models.Transaction, error)
	GetTransactionByID(transactionID, userID *uuid.UUID) (*models.Transaction, error)
	DeleteTransaction(transactionID uuid.UUID) error
	UpdateTransaction(transactionID uuid.UUID, transaction *models.Transaction) (*models.Transaction, error)
}
