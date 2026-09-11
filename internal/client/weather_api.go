package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Address  string `json:"address"`
	Timezone string `json:"timezone"`

	CurrentConditions struct {
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feelslike"`
		Conditions string  `json:"conditions"`
		Humidity   float64 `json:"humidity"`

		WindSpeed float64 `json:"windspeed"`
		WindDir   float64 `json:"winddir"`

		Precip     float64 `json:"precip"`
		PrecipProb float64 `json:"precipprob"`

		CloudCover float64 `json:"cloudcover"`
		Visibility float64 `json:"visibility"`
		Pressure   float64 `json:"pressure"`
		UVIndex    float64 `json:"uvindex"`

		Sunrise string `json:"sunrise"`
		Sunset  string `json:"sunset"`
		Icon    string `json:"icon"`
	} `json:"currentConditions"`
}

func GetWeather(city string) (WeatherResponse, error) {
	res, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + city + "?unitGroup=us&include=current&key=" + os.Getenv("WEATHER_API_KEY") + "&contentType=json")
	if err != nil {
		return WeatherResponse{}, errors.New("Something went wrong")
	}
	defer res.Body.Close()

	var data WeatherResponse

	fmt.Println("Response status code:", res.StatusCode)
	fmt.Println("Response body:", res.Body)
	fmt.Println("API KEY:", os.Getenv("WEATHER_API_KEY"))
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return WeatherResponse{}, errors.New("Error decoding JSON")
	}

	return data, nil

}
