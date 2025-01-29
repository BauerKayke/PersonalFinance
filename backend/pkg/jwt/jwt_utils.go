package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

// Chave secreta para assinar os tokens JWT (nunca exponha em código público)
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Claims representa os dados que vamos armazenar no token
type Claims struct {
	UserID *uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT gera um token JWT para o usuário
func GenerateJWT(userID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Expira em 24 horas
	claims := Claims{
		UserID: &userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT valida um token JWT e retorna as claims se válido
func ValidateJWT(tokenString string) (*uuid.UUID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}
	return nil, nil
}
