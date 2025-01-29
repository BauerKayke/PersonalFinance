package mappers

import (
	"backend/internal/dto"
	"backend/internal/models"
	"github.com/google/uuid"
	"time"
)

func ToCreditCardModel(dto dto.CreditCardCreateRequest, userID *uuid.UUID) models.CreditCard {
	return models.CreditCard{
		ID:            uuid.New(),
		UserID:        *userID,
		CardNumber:    dto.CardNumber,
		Expiration:    dto.Expiration,
		BankAccountID: dto.BankAccountID,
		CardName:      dto.CardName,
		CreditLimit:   dto.CreditLimit,
		Available:     dto.Available,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func ToCreditCardUpdateModel(dto dto.CreditCardUpdateRequest, userID *uuid.UUID) models.CreditCard {
	return models.CreditCard{
		UserID:        *userID,
		CardNumber:    dto.CardNumber,
		Expiration:    dto.Expiration,
		BankAccountID: dto.BankAccountID,
		CardName:      dto.CardName,
		CreditLimit:   dto.CreditLimit,
		Available:     dto.Available,
		UpdatedAt:     time.Now(),
	}
}
