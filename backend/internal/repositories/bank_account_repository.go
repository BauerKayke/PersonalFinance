package repositories

import (
	"backend/internal/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type BankAccountRepository struct {
	DB *sql.DB
}

func NewBankAccountRepository(db *sql.DB) *BankAccountRepository {
	return &BankAccountRepository{DB: db}
}

func (b *BankAccountRepository) CreateBankAccount(bankAccount *models.BankAccount) (*models.BankAccount, error) {
	bankAccount.ID = uuid.New()
	bankAccount.CreatedAt = time.Now()
	bankAccount.UpdatedAt = time.Now()
	query := `INSERT INTO bank_accounts (id, user_id, bank_name, account_no, balance, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`

	err := b.DB.QueryRow(query, bankAccount.ID, bankAccount.UserID, bankAccount.BankName, bankAccount.AccountNo,
		bankAccount.Balance, bankAccount.CreatedAt, bankAccount.UpdatedAt).Scan(&bankAccount.ID, &bankAccount.CreatedAt, &bankAccount.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not create bank account: %w", err)
	}

	return bankAccount, nil
}

func (b *BankAccountRepository) GetAllBankAccounts() ([]*models.BankAccount, error) {
	rows, err := b.DB.Query(`SELECT id, user_id, bank_name, account_no, balance, created_at, updated_at FROM bank_accounts`)
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

func (b *BankAccountRepository) GetBankAccountByID(id uuid.UUID) (*models.BankAccount, error) {
	query := `SELECT id, user_id, bank_name, account_no, balance, created_at, updated_at FROM bank_accounts WHERE id = $1`
	var bankAccount models.BankAccount
	err := b.DB.QueryRow(query, id).Scan(&bankAccount.ID, &bankAccount.UserID, &bankAccount.BankName, &bankAccount.AccountNo,
		&bankAccount.Balance, &bankAccount.CreatedAt, &bankAccount.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bank account not found: %w", err)
		}
		return nil, fmt.Errorf("could not get bank account by id: %w", err)
	}

	return &bankAccount, nil
}

func (b *BankAccountRepository) DeleteBankAccount(id uuid.UUID) error {
	query := `DELETE FROM bank_accounts WHERE id = $1`
	_, err := b.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete bank account: %w", err)
	}

	return nil
}

func (b *BankAccountRepository) UpdateBankAccount(id uuid.UUID, bankAccount *models.BankAccount) (*models.BankAccount, error) {
	query := `UPDATE bank_accounts SET bank_name = $1, account_no = $2, balance = $3, updated_at = $4 
	WHERE id = $5 RETURNING id, user_id, bank_name, account_no, balance, created_at, updated_at`

	err := b.DB.QueryRow(query, bankAccount.BankName, bankAccount.AccountNo, bankAccount.Balance,
		time.Now(), id).Scan(&bankAccount.ID, &bankAccount.UserID, &bankAccount.BankName, &bankAccount.AccountNo,
		&bankAccount.Balance, &bankAccount.CreatedAt, &bankAccount.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not update bank account: %w", err)
	}

	return bankAccount, nil
}
