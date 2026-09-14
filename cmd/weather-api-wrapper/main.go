package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/cache"
	"github.com/amanuel-tk/weather-api-wrapper/internal/config"
	"github.com/amanuel-tk/weather-api-wrapper/internal/handler"
	"github.com/amanuel-tk/weather-api-wrapper/internal/middleware"
	"github.com/amanuel-tk/weather-api-wrapper/internal/ratelimiter"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	ctx := context.Background()

	redisClient := cache.NewRedisClient(cfg.Redis.Address)

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	h := &handler.Handler{
		Redis:    redisClient,
		APIkey:   cfg.Weather.APIKey,
		CacheTTL: cfg.Redis.CacheTTL,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.GetWeather)

	r1 := ratelimiter.NewRateLimiter()
	r1.CleanUp()

	mw := &middleware.Middleware{
		RateLimiter: r1,
	}

	handleWithMiddleware := mw.Middleware(mux)

	http.ListenAndServe(":8080", handleWithMiddleware)
}
