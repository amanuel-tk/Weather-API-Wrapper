package handler

import (
	"errors"
	"fmt"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		errors.New("Something went wrong")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message:city is needed"))
	}

	fmt.Println(city)

}
