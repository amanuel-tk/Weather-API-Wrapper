package handler

import (
	"fmt"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Current weather informataion")

}
