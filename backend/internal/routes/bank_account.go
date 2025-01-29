package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type BankAccountRouter struct {
	*handlers.BankAccountHandler
}

func NewBankAccountRouter(handler *handlers.BankAccountHandler) *BankAccountRouter {
	return &BankAccountRouter{handler}
}

func (t *BankAccountRouter) RegisterRoutes(router chi.Router, authMiddleware *middleware.AuthMiddlewareHandler) {
	router.With(authMiddleware.AuthMiddleware).Route("/bank-account", func(r chi.Router) {
		r.Get("/", t.BankAccountHandler.GetAllBankAccount())
		r.Get("/{id}", t.BankAccountHandler.GetBankAccountByID())
		r.Post("/", t.BankAccountHandler.CreateBankAccount())
		r.Delete("/{id}", t.BankAccountHandler.DeleteBankAccount())
		r.Put("/{id}", t.BankAccountHandler.UpdateBankAccount())
	})
}
