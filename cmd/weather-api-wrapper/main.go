package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/cache"
	"github.com/amanuel-tk/weather-api-wrapper/internal/handler"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	ctx := context.Background()

	redisClient := cache.NewRedisClient()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	h := &handler.Handler{
		Redis: redisClient,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", h.GetWeather)

	http.ListenAndServe(":8080", mux)
}
