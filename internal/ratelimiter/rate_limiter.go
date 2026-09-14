package ratelimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Visitor struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}
type RateLimiterIP struct {
	ips map[string]*Visitor
	mu  sync.Mutex
}

func NewRateLimiter() *RateLimiterIP {
	return &RateLimiterIP{
		ips: make(map[string]*Visitor),
	}
}

func (i *RateLimiterIP) GetLimiter(ip string) *Visitor {

	i.mu.Lock()
	defer i.mu.Unlock()

	visitor, exists := i.ips[ip]

	if !exists {
		visitor = &Visitor{
			Limiter:  rate.NewLimiter(5, 10),
			LastSeen: time.Now(),
		}
		i.ips[ip] = visitor
	}

	return visitor
}

func (i *RateLimiterIP) CleanUp() {
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			i.mu.Lock()
			for ip, visitor := range i.ips {
				if time.Since(visitor.LastSeen) > time.Minute*10 {
					delete(i.ips, ip)
				}
			}
			i.mu.Unlock()
		}
	}()
}
