package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	var valueDecode interface{}

	fmt.Println("Request received", r.URL.Query())

	err := json.NewDecoder(r.Body).Decode(&valueDecode)

	if err != nil {
		errors.New("Something went wrong")
	}

	fmt.Println(valueDecode)

}
