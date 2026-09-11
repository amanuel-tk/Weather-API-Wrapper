package main

import (
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/handler"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.GetWeather)

	http.ListenAndServe(":8080", mux)
}
