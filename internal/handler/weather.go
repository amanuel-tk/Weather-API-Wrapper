package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/client"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	Redis *redis.Client
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "city is needed", http.StatusBadRequest)
		return
	}

	fmt.Println(city)

	cachedValue, err := h.Redis.Get(ctx, city).Result()

	if err == nil {
		fmt.Println(cachedValue)
	}
	if err != redis.Nil {
		fmt.Println("something is wrong with reddis", err)
	}

	data, err := client.GetWeather(city)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.Marshal(data)

	if err == nil {
		fmt.Println(jsonData)
		err = h.Redis.Set(ctx, city, jsonData, 0).Err()

	}

	fmt.Println(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
