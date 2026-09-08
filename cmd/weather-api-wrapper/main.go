package main

import (
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.GetWeather)

	http.ListenAndServe(":8080", mux)
}
