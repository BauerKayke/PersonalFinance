package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type CreditCardRepositories interface {
	CreateCreditCard(creditCard *models.CreditCard) (*models.CreditCard, error)
	GetAllCreditCards() ([]*models.CreditCard, error)
	GetCreditCardByID(id uuid.UUID) (*models.CreditCard, error)
	DeleteCreditCard(id uuid.UUID) error
	UpdateCreditCard(id uuid.UUID, creditCard *models.CreditCard) (*models.CreditCard, error)
}
