package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
		var data client.WeatherResponse

		if err := json.Unmarshal([]byte(cachedValue), &data); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return
		}
		h.Redis.Del(ctx, city)
		fmt.Println("bad cache, fetching fresh data")
	} else if err != redis.Nil {
		fmt.Println("something is wrong with redis", err)
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

		err = h.Redis.Set(ctx, city, jsonData, 10*time.Second).Err()

	}

	fmt.Println(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
