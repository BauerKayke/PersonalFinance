package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"net/http"
	"time"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	limiter := tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpError := tollbooth.LimitByRequest(limiter, w, r)
		if httpError != nil {
			w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}
		next.ServeHTTP(w, r)
	})
}
