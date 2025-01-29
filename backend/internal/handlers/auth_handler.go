package handlers

import (
	"backend/internal/interfaces"
	"backend/internal/models"
	"backend/pkg/jwt"
	"github.com/bootcamp-go/web/response"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

type AuthHandler struct {
	interfaces.AuthServices
}

func NewAuthHandler(auth interfaces.AuthServices) *AuthHandler {
	return &AuthHandler{auth}
}

// LoginHandler gera o token JWT e salva a sessão no Redis
func (a *AuthHandler) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := a.AuthServices.GetUserByEmail(email)
		if err != nil || user.Password != password {
			response.Error(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		token, err := jwt.GenerateJWT(user.ID)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		session := &models.Sessions{
			ID:        uuid.New(),
			UserID:    user.ID,
			Token:     token,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ok, err := a.AuthServices.SaveSession(session)
		if err != nil || !ok {
			response.Error(w, http.StatusInternalServerError, "Failed to save session")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
		response.JSON(w, http.StatusOK, "Login successful")
	}
}

// LogoutHandler remove a sessão do Redis e adiciona à blacklist
func (a *AuthHandler) LogoutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			response.Error(w, http.StatusUnauthorized, "Token is missing")
			return
		}

		parts := strings.Split(tokenHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(w, http.StatusUnauthorized, "Invalid token format")
			return
		}

		token := parts[1]
		_, err := jwt.ValidateJWT(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized: "+err.Error())
			return
		}

		if err := a.AuthServices.DeleteSession(token); err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to delete session")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})
		response.JSON(w, http.StatusOK, "Logout successful")

	}
}
