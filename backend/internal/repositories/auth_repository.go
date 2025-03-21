package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (a AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	var user models.User

	err := a.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (a AuthRepository) SaveSession(session *models.Sessions) (bool, error) {
	ok, err := a.GetActiveSession(&session.UserID)
	if err != nil {
		return false, err
	}
	if ok {
		return false, errors.New("session already active")
	}
	query := `INSERT INTO sessions (id, user_id, token, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = a.DB.Exec(query, session.ID, session.UserID, session.Token, session.IsActive, session.CreatedAt, session.UpdatedAt)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a AuthRepository) DeleteSession(userId *uuid.UUID) error {
	ok, err := a.GetActiveSession(userId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("session not found")
	}
	query := `DELETE FROM sessions WHERE user_id = $1`
	_, err = a.DB.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (a AuthRepository) GetActiveSession(userId *uuid.UUID) (bool, error) {
	query := `SELECT id, user_id, token, is_active FROM sessions WHERE user_id = $1`
	var session models.Sessions

	err := a.DB.QueryRow(query, userId).Scan(&session.ID, &session.UserID, &session.Token, &session.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return session.IsActive, nil
		}
		return session.IsActive, err
	}

	return session.IsActive, nil
}
