package handler

import (
	"fmt"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message:city is needed"))
	}

	fmt.Println(city)

}
