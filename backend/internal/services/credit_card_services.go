package services

import (
	"backend/internal/interfaces"
	"backend/internal/models"
	"github.com/google/uuid"
)

type CreditCardService struct {
	Repo interfaces.CreditCardRepositories
}

func NewCreditCardService(repo interfaces.CreditCardRepositories) interfaces.CreditCardServices {
	return &CreditCardService{Repo: repo}
}

func (c CreditCardService) CreateCreditCard(creditCard *models.CreditCard) (*models.CreditCard, error) {
	return c.Repo.CreateCreditCard(creditCard)
}

func (c CreditCardService) GetAllCreditCards(userId *uuid.UUID) ([]*models.CreditCard, error) {
	return c.Repo.GetAllCreditCards(userId)
}

func (c CreditCardService) GetCreditCardByID(creditCardId, userId *uuid.UUID) (*models.CreditCard, error) {
	return c.Repo.GetCreditCardByID(creditCardId, userId)
}

func (c CreditCardService) DeleteCreditCard(creditCardId uuid.UUID) error {
	return c.Repo.DeleteCreditCard(creditCardId)
}

func (c CreditCardService) UpdateCreditCard(creditCardId uuid.UUID, creditCard *models.CreditCard) (*models.CreditCard, error) {
	return c.UpdateCreditCard(creditCardId, creditCard)
}
