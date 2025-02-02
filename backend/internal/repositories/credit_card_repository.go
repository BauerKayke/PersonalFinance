package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreditCardRepository struct {
	DB              *sql.DB
	UserRepo        *UserRepository
	BankAccountRepo *BankAccountRepository
}

func NewCreditCardRepository(db *sql.DB, userRepo *UserRepository, bankAccountRepo *BankAccountRepository) *CreditCardRepository {
	return &CreditCardRepository{
		DB:              db,
		UserRepo:        userRepo,
		BankAccountRepo: bankAccountRepo,
	}
}

func (c *CreditCardRepository) CreateCreditCard(creditCard *models.CreditCard) (*models.CreditCard, error) {
	userExists, err := c.UserRepo.UserExists(creditCard.UserID)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence: %w", err)
	}
	if !userExists {
		return nil, errors.New("user does not exist")
	}

	accountExists, err := c.BankAccountRepo.BankAccountExist(creditCard.BankAccountID)
	if err != nil {
		return nil, fmt.Errorf("error checking bank account existence: %w", err)
	}
	if !accountExists {
		return nil, errors.New("bank account does not exist")
	}

	creditCard.ID = uuid.New()
	creditCard.CreatedAt = time.Now()
	creditCard.UpdatedAt = time.Now()

	query := `INSERT INTO credit_cards (id, user_id, bank_account_id, card_name, card_number, credit_limit, available, expiration, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = c.DB.Exec(query,
		creditCard.ID,
		creditCard.UserID,
		creditCard.BankAccountID,
		creditCard.CardName,
		creditCard.CardNumber,
		creditCard.CreditLimit,
		creditCard.Available,
		creditCard.Expiration,
		creditCard.CreatedAt,
		creditCard.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("could not create credit card: %w", err)
	}

	return creditCard, nil
}

func (c *CreditCardRepository) UpdateCreditCard(id uuid.UUID, creditCard *models.CreditCard) (*models.CreditCard, error) {
	// Verificar existência do cartão
	exists, err := c.CreditCardExists(id)
	if err != nil {
		return nil, fmt.Errorf("error checking credit card existence: %w", err)
	}
	if !exists {
		return nil, errors.New("credit card does not exist")
	}

	query := `UPDATE credit_cards 
	SET card_name = $1, card_number = $2, credit_limit = $3, available = $4, expiration = $5, updated_at = $6 
	WHERE id = $7`

	result, err := c.DB.Exec(query,
		creditCard.CardName,
		creditCard.CardNumber,
		creditCard.CreditLimit,
		creditCard.Available,
		creditCard.Expiration,
		time.Now(),
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("could not update credit card: %w", err)
	}

	// Verificar se a atualização foi efetiva
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, errors.New("no credit card updated")
	}

	// Retornar o cartão atualizado
	return c.GetCreditCardByID(&id, &creditCard.UserID)
}

func (c *CreditCardRepository) DeleteCreditCard(id uuid.UUID) error {
	exists, err := c.CreditCardExists(id)
	if err != nil {
		return fmt.Errorf("error checking credit card existence: %w", err)
	}
	if !exists {
		return errors.New("credit card does not exist")
	}

	result, err := c.DB.Exec(`DELETE FROM credit_cards WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("could not delete credit card: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no credit card deleted")
	}

	return nil
}

func (c *CreditCardRepository) CreditCardExists(id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM credit_cards WHERE id = $1)`
	var exists bool
	err := c.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking credit card existence: %w", err)
	}
	return exists, nil
}

func (c *CreditCardRepository) GetAllCreditCards(userId *uuid.UUID) ([]*models.CreditCard, error) {
	query := `SELECT id, user_id, bank_account_id, card_name, card_number, credit_limit, available, expiration, created_at, updated_at 
	FROM credit_cards WHERE user_id = $1`

	rows, err := c.DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("could not get credit cards: %w", err)
	}
	defer rows.Close()

	var creditCards []*models.CreditCard
	for rows.Next() {
		var cc models.CreditCard
		err := rows.Scan(
			&cc.ID,
			&cc.UserID,
			&cc.BankAccountID,
			&cc.CardName,
			&cc.CardNumber,
			&cc.CreditLimit,
			&cc.Available,
			&cc.Expiration,
			&cc.CreatedAt,
			&cc.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan credit card: %w", err)
		}
		creditCards = append(creditCards, &cc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return creditCards, nil
}

func (c *CreditCardRepository) GetCreditCardByID(id, userId *uuid.UUID) (*models.CreditCard, error) {
	query := `SELECT id, user_id, bank_account_id, card_name, card_number, credit_limit, available, expiration, created_at, updated_at 
	FROM credit_cards WHERE id = $1 AND user_id = $2`

	var cc models.CreditCard
	err := c.DB.QueryRow(query, id, userId).Scan(
		&cc.ID,
		&cc.UserID,
		&cc.BankAccountID,
		&cc.CardName,
		&cc.CardNumber,
		&cc.CreditLimit,
		&cc.Available,
		&cc.Expiration,
		&cc.CreatedAt,
		&cc.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("credit card not found: %w", err)
		}
		return nil, fmt.Errorf("could not get credit card: %w", err)
	}

	return &cc, nil
}
