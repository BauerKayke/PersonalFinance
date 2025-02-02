package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransactionRepository struct {
	DB              *sql.DB
	UserRepo        *UserRepository
	BankAccountRepo *BankAccountRepository
	CreditCardRepo  *CreditCardRepository
}

func NewTransactionRepository(
	db *sql.DB,
	userRepo *UserRepository,
	bankAccountRepo *BankAccountRepository,
	creditCardRepo *CreditCardRepository,
) *TransactionRepository {
	return &TransactionRepository{
		DB:              db,
		UserRepo:        userRepo,
		BankAccountRepo: bankAccountRepo,
		CreditCardRepo:  creditCardRepo,
	}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	// Validações iniciais
	if transaction.BankAccountID == nil && transaction.CreditCardID == nil {
		return nil, errors.New("transaction must be linked to a bank account or credit card")
	}

	// Verificar existência do usuário
	userExists, err := r.UserRepo.UserExists(transaction.UserID)
	if err != nil {
		return nil, fmt.Errorf("error verifying user: %w", err)
	}
	if !userExists {
		return nil, errors.New("user does not exist")
	}

	// Verificar conta bancária se informada
	if transaction.BankAccountID != nil {
		account, err := r.BankAccountRepo.GetBankAccountByID(transaction.BankAccountID, &transaction.UserID)
		if err != nil {
			return nil, fmt.Errorf("error verifying bank account: %w", err)
		}
		if account == nil || account.UserID != transaction.UserID {
			return nil, errors.New("invalid bank account for user")
		}
	}

	// Verificar cartão de crédito se informado
	if transaction.CreditCardID != nil {
		card, err := r.CreditCardRepo.GetCreditCardByID(transaction.CreditCardID, &transaction.UserID)
		if err != nil {
			return nil, fmt.Errorf("error verifying credit card: %w", err)
		}
		if card == nil || card.UserID != transaction.UserID {
			return nil, errors.New("invalid credit card for user")
		}
	}

	// Configurar campos gerados
	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	// Tratar descrição opcional
	var description sql.NullString
	if transaction.Description != nil {
		description = sql.NullString{String: *transaction.Description, Valid: true}
	}

	query := `INSERT INTO transactions (
		id, user_id, bank_account_id, credit_card_id, amount, 
		category, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = r.DB.Exec(query,
		transaction.ID,
		transaction.UserID,
		transaction.BankAccountID,
		transaction.CreditCardID,
		transaction.Amount,
		transaction.Category,
		description,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(transactionId uuid.UUID, transaction *models.Transaction) (*models.Transaction, error) {
	// Verificar existência da transação
	exists, err := r.TransactionExists(transactionId)
	if err != nil {
		return nil, fmt.Errorf("error checking transaction: %w", err)
	}
	if !exists {
		return nil, errors.New("transaction not found")
	}

	// Obter transação existente
	current, err := r.GetTransactionByID(&transactionId, &transaction.UserID)
	if err != nil {
		return nil, fmt.Errorf("error fetching transaction: %w", err)
	}

	// Validar novos relacionamentos
	if transaction.BankAccountID != nil {
		account, err := r.BankAccountRepo.GetBankAccountByID(transaction.BankAccountID, &transaction.UserID)
		if err != nil || account == nil || account.UserID != current.UserID {
			return nil, errors.New("invalid bank account for transaction update")
		}
	}

	if transaction.CreditCardID != nil {
		card, err := r.CreditCardRepo.GetCreditCardByID(transaction.CreditCardID, &transaction.UserID)
		if err != nil || card == nil || card.UserID != current.UserID {
			return nil, errors.New("invalid credit card for transaction update")
		}
	}

	// Atualizar campos
	transaction.UpdatedAt = time.Now()
	var description sql.NullString
	if transaction.Description != nil {
		description = sql.NullString{String: *transaction.Description, Valid: true}
	}

	query := `UPDATE transactions SET
		bank_account_id = $1,
		credit_card_id = $2,
		amount = $3,
		category = $4,
		description = $5,
		updated_at = $6
	WHERE id = $7`

	_, err = r.DB.Exec(query,
		transaction.BankAccountID,
		transaction.CreditCardID,
		transaction.Amount,
		transaction.Category,
		description,
		transaction.UpdatedAt,
		transactionId,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	// Retornar transação atualizada
	return r.GetTransactionByID(&transactionId, &transaction.UserID)
}

func (r *TransactionRepository) DeleteTransaction(id uuid.UUID) error {
	exists, err := r.TransactionExists(id)
	if err != nil {
		return fmt.Errorf("error checking transaction: %w", err)
	}
	if !exists {
		return errors.New("transaction not found")
	}

	result, err := r.DB.Exec("DELETE FROM transactions WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete transaction: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verifying deletion: %w", err)
	}

	if rows == 0 {
		return errors.New("no transaction deleted")
	}

	return nil
}

func (r *TransactionRepository) TransactionExists(id uuid.UUID) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM transactions WHERE id = $1)"
	var exists bool
	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking transaction existence: %w", err)
	}
	return exists, nil
}

func (r *TransactionRepository) GetAllTransactions(userID *uuid.UUID) ([]*models.Transaction, error) {
	query := `SELECT id, user_id, bank_account_id, credit_card_id, amount, category, description, created_at, updated_at FROM transactions WHERE user_id = $1`
	rows, err := r.DB.Query(query, &userID)
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

func (r *TransactionRepository) GetTransactionByID(transactionID, userID *uuid.UUID) (*models.Transaction, error) {
	query := `SELECT id, user_id, bank_account_id, credit_card_id, amount, category, description, created_at, updated_at 
			  FROM transactions WHERE id = $1 AND user_id = $2`
	var transaction models.Transaction
	var description sql.NullString

	err := r.DB.QueryRow(query, transactionID, userID).Scan(&transaction.ID, &transaction.UserID, &transaction.BankAccountID, &transaction.CreditCardID,
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
