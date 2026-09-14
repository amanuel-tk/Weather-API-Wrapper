package middleware

import (
	"net"
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/ratelimiter"
)

type Middleware struct {
	RateLimiter *ratelimiter.RateLimiterIP
}

func (i *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid client address", http.StatusBadRequest)
			return
		}

		limiter := i.RateLimiter.GetLimiter(ip)

		if !limiter.Limiter.Allow() {
			http.Error(w, "Too many request", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
