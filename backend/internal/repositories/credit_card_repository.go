package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreditCardRepositories struct {
	DB *sql.DB
}

func NewCreditCardRepository(db *sql.DB) *CreditCardRepositories {
	return &CreditCardRepositories{DB: db}
}

func (c CreditCardRepositories) CreateCreditCard(creditCard *models.CreditCard) (*models.CreditCard, error) {
	query := `INSERT INTO credit_cards (id, user_id, bank_account_id, card_name, card_number, credit_limit, available, expiration, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	creditCard.ID = uuid.New()
	creditCard.CreatedAt = time.Now()
	creditCard.UpdatedAt = time.Now()

	_, err := c.DB.Exec(query, creditCard.ID, creditCard.UserID, creditCard.BankAccountID, creditCard.CardName, creditCard.CardNumber, creditCard.CreditLimit, creditCard.Available, creditCard.Expiration, creditCard.CreatedAt, creditCard.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return creditCard, nil
}

func (c CreditCardRepositories) GetAllCreditCards() ([]*models.CreditCard, error) {
	query := `SELECT id, user_id, bank_account_id, card_name, card_number, credit_limit, available, expiration, created_at, updated_at FROM credit_cards`
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var creditCards []*models.CreditCard
	for rows.Next() {
		var creditCard models.CreditCard
		if err := rows.Scan(&creditCard.ID, &creditCard.UserID, &creditCard.BankAccountID, &creditCard.CardName, &creditCard.CardNumber, &creditCard.CreditLimit, &creditCard.Available, &creditCard.Expiration, &creditCard.CreatedAt, &creditCard.UpdatedAt); err != nil {
			return nil, err
		}
		creditCards = append(creditCards, &creditCard)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return creditCards, nil
}

func (c CreditCardRepositories) GetCreditCardByID(id uuid.UUID) (*models.CreditCard, error) {
	query := `SELECT id, user_id, bank_account_id, card_name, card_number, credit_limit, available, expiration, created_at, updated_at FROM credit_cards WHERE id = $1`
	var creditCard models.CreditCard

	err := c.DB.QueryRow(query, id).Scan(&creditCard.ID, &creditCard.UserID, &creditCard.BankAccountID, &creditCard.CardName, &creditCard.CardNumber, &creditCard.CreditLimit, &creditCard.Available, &creditCard.Expiration, &creditCard.CreatedAt, &creditCard.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &creditCard, nil
}

func (c CreditCardRepositories) DeleteCreditCard(id uuid.UUID) error {
	query := `DELETE FROM credit_cards WHERE id = $1`
	result, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no credit card found with the given ID")
	}

	return nil
}

func (c CreditCardRepositories) UpdateCreditCard(id uuid.UUID, creditCard *models.CreditCard) (*models.CreditCard, error) {
	query := `UPDATE credit_cards SET card_name = $1, card_number = $2, credit_limit = $3, available = $4, expiration = $5, updated_at = $6 WHERE id = $7`
	creditCard.UpdatedAt = time.Now()

	result, err := c.DB.Exec(query, creditCard.CardName, creditCard.CardNumber, creditCard.CreditLimit, creditCard.Available, creditCard.Expiration, creditCard.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no credit card found with the given ID")
	}

	creditCard.ID = uuid.MustParse(fmt.Sprintf("%d", id)) // Ajuste se necessário.
	return creditCard, nil
}
