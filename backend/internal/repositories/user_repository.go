package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u UserRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6)`
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.DB.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1`
	var user models.User

	err := u.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	log.Println("User found")
	log.Println(user)
	return &user, nil
}

func (u UserRepository) GetAllUsers() ([]*models.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users`
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`
	var user models.User

	err := u.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := u.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

func (u UserRepository) UpdateUser(id uuid.UUID, user *models.User) (*models.User, error) {
	query := `UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5`
	user.UpdatedAt = time.Now()

	result, err := u.DB.Exec(query, user.Name, user.Email, user.Password, user.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user found with the given ID")
	}

	return user, nil
}
