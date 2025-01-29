package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type CreditCardRouter struct {
	*handlers.CreditCardHandler
}

func NewCreditCardRouter(handler *handlers.CreditCardHandler) *CreditCardRouter {
	return &CreditCardRouter{handler}
}

func (t *CreditCardRouter) RegisterRoutes(router chi.Router, authMiddleware *middleware.AuthMiddlewareHandler) {
	router.With(authMiddleware.AuthMiddleware).Route("/credit-card", func(r chi.Router) {
		r.Get("/", t.CreditCardHandler.GetAllCreditCard())
		r.Get("/{id}", t.CreditCardHandler.GetCreditCardByID())
		r.Post("/", t.CreditCardHandler.CreateCreditCard())
		r.Delete("/{id}", t.CreditCardHandler.DeleteCreditCard())
		r.Put("/{id}", t.CreditCardHandler.UpdateCreditCard())
	})
}
