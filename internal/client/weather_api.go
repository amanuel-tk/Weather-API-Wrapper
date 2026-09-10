package client

import (
	"encoding/json"
	"errors"
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

func GetWeather(city string) (WeatherResponse, error) {
	res, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + city + "?unitGroup=us&include=current&key=EN3UGBJT66YSWVX4HR3MT5KDU&contentType=json")
	if err != nil {
		return WeatherResponse{}, errors.New("Something went wrong")
	}
	defer res.Body.Close()

	var data WeatherResponse

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return WeatherResponse{}, errors.New("Error decoding JSON")
	}

	return data, nil

}
