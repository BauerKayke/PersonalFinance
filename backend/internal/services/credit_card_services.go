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

func (c CreditCardService) GetAllCreditCards() ([]*models.CreditCard, error) {
	return c.Repo.GetAllCreditCards()
}

func (c CreditCardService) GetCreditCardByID(id uuid.UUID) (*models.CreditCard, error) {
	return c.Repo.GetCreditCardByID(id)
}

func (c CreditCardService) DeleteCreditCard(id uuid.UUID) error {
	return c.Repo.DeleteCreditCard(id)
}

func (c CreditCardService) UpdateCreditCard(id uuid.UUID, creditCard *models.CreditCard) (*models.CreditCard, error) {
	return c.UpdateCreditCard(id, creditCard)
}
