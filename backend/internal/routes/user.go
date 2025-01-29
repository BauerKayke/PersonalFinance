package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	*handlers.UserHandler
}

func NewUserRouter(handler *handlers.UserHandler) *UserRouter {
	return &UserRouter{handler}
}

func (t *UserRouter) RegisterRoutes(router chi.Router, authMiddleware *middleware.AuthMiddlewareHandler) {
	router.Route("/user", func(r chi.Router) {
		r.With(authMiddleware.AuthMiddleware).Get("/", t.UserHandler.GetAllUser())
		r.With(authMiddleware.AuthMiddleware).Get("/{id}", t.UserHandler.GetUserByID())
		r.Post("/", t.UserHandler.CreateUser())
		r.With(authMiddleware.AuthMiddleware).Delete("/{id}", t.UserHandler.DeleteUser())
		r.With(authMiddleware.AuthMiddleware).Put("/{id}", t.UserHandler.UpdateUser())
	})
}
