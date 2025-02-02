package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BankAccountRepository struct {
	DB *sql.DB
	// Adicione uma referência ao UserRepository se necessário verificar usuários
	UserRepo *UserRepository
}

func NewBankAccountRepository(db *sql.DB, userRepo *UserRepository) *BankAccountRepository {
	return &BankAccountRepository{
		DB:       db,
		UserRepo: userRepo,
	}
}

func (b *BankAccountRepository) CreateBankAccount(bankAccount *models.BankAccount) (*models.BankAccount, error) {
	// Verificar se o usuário existe
	exists, err := b.UserRepo.UserExists(bankAccount.UserID)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence: %w", err)
	}
	if !exists {
		return nil, errors.New("user does not exist")
	}

	bankAccount.ID = uuid.New()
	bankAccount.CreatedAt = time.Now()
	bankAccount.UpdatedAt = time.Now()

	query := `INSERT INTO bank_accounts (id, user_id, bank_name, account_no, balance, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = b.DB.Exec(query,
		bankAccount.ID,
		bankAccount.UserID,
		bankAccount.BankName,
		bankAccount.AccountNo,
		bankAccount.Balance,
		bankAccount.CreatedAt,
		bankAccount.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("could not create bank account: %w", err)
	}

	return bankAccount, nil
}

func (b *BankAccountRepository) UpdateBankAccount(id uuid.UUID, bankAccount *models.BankAccount) (*models.BankAccount, error) {
	// Verificar existência
	exists, err := b.BankAccountExist(&id)
	if err != nil {
		return nil, fmt.Errorf("error checking account existence: %w", err)
	}
	if !exists {
		return nil, errors.New("bank account does not exist")
	}

	// Atualizar com verificação de linhas afetadas
	query := `UPDATE bank_accounts 
	SET bank_name = $1, account_no = $2, balance = $3, updated_at = $4 
	WHERE id = $5`

	result, err := b.DB.Exec(query,
		bankAccount.BankName,
		bankAccount.AccountNo,
		bankAccount.Balance,
		time.Now(),
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("could not update bank account: %w", err)
	}

	// Verificar se realmente atualizou
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, errors.New("no bank account updated")
	}

	// Buscar versão atualizada
	return b.GetBankAccountByID(&id, &bankAccount.UserID)
}

func (b *BankAccountRepository) DeleteBankAccount(id uuid.UUID) error {
	exists, err := b.BankAccountExist(&id)
	if err != nil {
		return fmt.Errorf("error checking account existence: %w", err)
	}
	if !exists {
		return errors.New("bank account does not exist")
	}

	result, err := b.DB.Exec(`DELETE FROM bank_accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("could not delete bank account: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no bank account deleted")
	}

	return nil
}

func (b *BankAccountRepository) BankAccountExist(id *uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM bank_accounts WHERE id = $1)`
	var exists bool
	err := b.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking account existence: %w", err)
	}
	return exists, nil
}

// Mantenha os outros métodos como GetAllBankAccounts e GetBankAccountByID com tratamento similar

func (b *BankAccountRepository) GetAllBankAccounts(userId *uuid.UUID) ([]*models.BankAccount, error) {
	rows, err := b.DB.Query(`SELECT id, user_id, bank_name, account_no, balance, created_at, updated_at FROM bank_accounts WHERE user_id = $1`, userId)
	if err != nil {
		return nil, fmt.Errorf("could not get all bank accounts: %w", err)
	}
	defer rows.Close()

	var bankAccounts []*models.BankAccount
	for rows.Next() {
		var bankAccount models.BankAccount
		if err := rows.Scan(&bankAccount.ID, &bankAccount.UserID, &bankAccount.BankName, &bankAccount.AccountNo,
			&bankAccount.Balance, &bankAccount.CreatedAt, &bankAccount.UpdatedAt); err != nil {
			return nil, fmt.Errorf("could not scan bank account: %w", err)
		}
		bankAccounts = append(bankAccounts, &bankAccount)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return bankAccounts, nil
}

func (b *BankAccountRepository) GetBankAccountByID(bankAccountId, userId *uuid.UUID) (*models.BankAccount, error) {
	query := `SELECT id, user_id, bank_name, account_no, balance, created_at, updated_at FROM bank_accounts WHERE id = $1 AND user_id = $2`
	var bankAccount models.BankAccount
	err := b.DB.QueryRow(query, bankAccountId, userId).Scan(&bankAccount.ID, &bankAccount.UserID, &bankAccount.BankName, &bankAccount.AccountNo,
		&bankAccount.Balance, &bankAccount.CreatedAt, &bankAccount.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bank account not found: %w", err)
		}
		return nil, fmt.Errorf("could not get bank account by id: %w", err)
	}

	return &bankAccount, nil
}
