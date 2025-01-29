package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	query := `INSERT INTO transactions (user_id, bank_account_id, credit_card_id, amount, category, description, created_at, updated_at) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	description := ""
	if transaction.Description != nil {
		description = *transaction.Description
	}

	_, err := r.DB.Exec(query, transaction.UserID, transaction.BankAccountID, transaction.CreditCardID,
		transaction.Amount, transaction.Category, description, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) GetAllTransactions() ([]*models.Transaction, error) {
	query := `SELECT id, user_id, bank_account_id, credit_card_id, amount, category, description, created_at, updated_at FROM transactions`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		var description sql.NullString
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.BankAccountID, &transaction.CreditCardID,
			&transaction.Amount, &transaction.Category, &description, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, err
		}
		if description.Valid {
			descriptionValue := description.String
			transaction.Description = &descriptionValue
		}
		transactions = append(transactions, &transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	query := `SELECT id, user_id, bank_account_id, credit_card_id, amount, category, description, created_at, updated_at 
			  FROM transactions WHERE id = $1`
	var transaction models.Transaction
	var description sql.NullString

	err := r.DB.QueryRow(query, id).Scan(&transaction.ID, &transaction.UserID, &transaction.BankAccountID, &transaction.CreditCardID,
		&transaction.Amount, &transaction.Category, &description, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if description.Valid {
		descriptionValue := description.String
		transaction.Description = &descriptionValue
	}

	return &transaction, nil
}

func (r *TransactionRepository) DeleteTransaction(id uuid.UUID) error {
	query := `DELETE FROM transactions WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no transaction found with the given ID")
	}

	return nil
}

func (r *TransactionRepository) UpdateTransaction(id uuid.UUID, transaction *models.Transaction) (*models.Transaction, error) {
	query := `UPDATE transactions SET user_id = $1, bank_account_id = $2, credit_card_id = $3, amount = $4, category = $5,
   description = $6, updated_at = $7 WHERE id = $8`
	description := ""
	if transaction.Description != nil {
		description = *transaction.Description
	}

	result, err := r.DB.Exec(query, transaction.UserID, transaction.BankAccountID, transaction.CreditCardID,
		transaction.Amount, transaction.Category, description, time.Now(), id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no transaction found with the given ID")
	}

	transaction.UpdatedAt = time.Now()
	return transaction, nil
}
