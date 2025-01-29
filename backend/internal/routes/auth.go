package routes

import (
	"backend/internal/handlers"
	"github.com/go-chi/chi/v5"
)

// AuthRouter define as rotas de autenticação
type AuthRouter struct {
	*handlers.AuthHandler
}

// NewAuthRouter cria um novo AuthRouter
func NewAuthRouter(handler *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{handler}
}

// RegisterRoutes registra as rotas de autenticação
func (a *AuthRouter) RegisterRoutes(router chi.Router) {
	// Rota de login (gera o token JWT)
	router.Post("/login", a.AuthHandler.LoginUser())
	router.Post("/logout", a.AuthHandler.LogoutUser())
}
