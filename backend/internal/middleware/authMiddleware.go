package middleware

import (
	"backend/internal/interfaces"
	"backend/pkg/jwt"
	"context"
	"github.com/bootcamp-go/web/response"
	"net/http"
)

type AuthMiddlewareHandler struct {
	authService interfaces.AuthServices
}

func NewAuthMiddleware(authService interfaces.AuthServices) *AuthMiddlewareHandler {
	return &AuthMiddlewareHandler{authService: authService}
}

func (a *AuthMiddlewareHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized: Token is missing", http.StatusUnauthorized)
			return
		}
		token := cookie.Value
		claims, err := jwt.ValidateJWT(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized: "+err.Error())
			return
		}

		ok, err := a.authService.GetActiveSession(claims)
		if err != nil || !ok {
			response.Error(w, http.StatusUnauthorized, "Unauthorized: Session not found or inactive")
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
