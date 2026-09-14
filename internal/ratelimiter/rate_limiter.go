package ratelimiter

import (
	"sync"

	"golang.org/x/time/rate"
)

type RateLimiterIP struct {
	ips map[string]*rate.Limiter
	mu  sync.Mutex
}

func NewRateLimiter() *RateLimiterIP {
	return &RateLimiterIP{
		ips: make(map[string]*rate.Limiter),
	}
}

func (i *RateLimiterIP) GetLimiter(ip string) *rate.Limiter {

	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]

	if !exists {
		limiter = rate.NewLimiter(5, 10)
		i.ips[ip] = limiter
	}

	return limiter
}
