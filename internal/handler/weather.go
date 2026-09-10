package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amanuel-tk/weather-api-wrapper/internal/client"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message:city is needed"))
	}

	fmt.Println(city)

	data, err := client.GetWeather(city)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message:" + err.Error()))

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
