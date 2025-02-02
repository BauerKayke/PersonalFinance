package interfaces

import (
	"backend/internal/models"
	"github.com/google/uuid"
)

type CreditCardRepositories interface {
	CreateCreditCard(creditCard *models.CreditCard) (*models.CreditCard, error)
	GetAllCreditCards(userID *uuid.UUID) ([]*models.CreditCard, error)
	GetCreditCardByID(creditCardId, userID *uuid.UUID) (*models.CreditCard, error)
	DeleteCreditCard(creditCardId uuid.UUID) error
	UpdateCreditCard(creditCardId uuid.UUID, creditCard *models.CreditCard) (*models.CreditCard, error)
}
