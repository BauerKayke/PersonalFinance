package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type TransactionRouter struct {
	*handlers.TransactionHandler
}

func NewTransactionRouter(handler *handlers.TransactionHandler) *TransactionRouter {
	return &TransactionRouter{handler}
}

func (t *TransactionRouter) RegisterRoutes(router chi.Router, authMiddleware *middleware.AuthMiddlewareHandler) {
	router.With(authMiddleware.AuthMiddleware).Route("/transactions", func(r chi.Router) {
		r.Get("/", t.TransactionHandler.GetAllTransactions())
		r.Get("/{id}", t.TransactionHandler.GetTransactionByID())
		r.Post("/", t.TransactionHandler.CreateTransaction())
		r.Delete("/{id}", t.TransactionHandler.DeleteTransaction())
		r.Put("/{id}", t.TransactionHandler.UpdateTransaction())
	})
}
