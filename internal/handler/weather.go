package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Location                 string  `json:"location"`
	Timezone                 string  `json:"timezone"`
	Temperature              float64 `json:"temperature"`
	FeelsLike                float64 `json:"feels_like"`
	Condition                string  `json:"condition"`
	Humidity                 float64 `json:"humidity"`
	WindSpeed                float64 `json:"wind_speed"`
	WindDirection            float64 `json:"wind_direction"`
	Precipitation            float64 `json:"precipitation"`
	PrecipitationProbability float64 `json:"precipitation_probability"`
	CloudCover               float64 `json:"cloud_cover"`
	Visibility               float64 `json:"visibility"`
	Pressure                 float64 `json:"pressure"`
	UVIndex                  float64 `json:"uv_index"`
	Sunrise                  string  `json:"sunrise"`
	Sunset                   string  `json:"sunset"`
	Icon                     string  `json:"icon"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		errors.New("Something went wrong")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message:city is needed"))
	}

	fmt.Println(city)

	res, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/addis%20ababa?unitGroup=us&include=current&key=EN3UGBJT66YSWVX4HR3MT5KDU&contentType=json")
	if err != nil {
		errors.New("Something went wrong")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message:failed to fetch weather data"))
		return
	}
	defer res.Body.Close()

	var newData any

	err = json.NewDecoder(res.Body).Decode(&newData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	pretty, err := json.MarshalIndent(newData, "", " ")

	fmt.Println(string(pretty))

	w.WriteHeader(http.StatusFound)
	w.Write([]byte(string(pretty)))

}
