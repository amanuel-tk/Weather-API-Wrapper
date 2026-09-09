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

	res, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/addis%20ababa?unitGroup=us&key=EN3UGBJT66YSWVX4HR3MT5KDU&contentType=json")
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message:failed to fetch weather data"))
		return
	}
	defer res.Body.Close()

	fmt.Println(res.Body)

}
