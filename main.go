package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rainycape/unidecode"
)

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type Current struct {
	LastUpdatedEpoch int       `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	UV               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type RequestData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		panic(".env does not load")
	}

	var input string

	fmt.Print("What is the city name? ")
	fmt.Scanln(&input)

	// convert to unicode standart
	input = unidecode.Unidecode(input)

	url := "https://api.weatherapi.com/v1/current.json?q=" + input + "&lang=pt&key=" + os.Getenv("WEATHER_API_KEY")

	response, err := http.Get(url)

	if err != nil {
		panic("We have a problem with the request")
	}

	var locationData RequestData

	if err := json.NewDecoder(response.Body).Decode(&locationData); err != nil {
		panic("The convertion type has an error")
	}

	if locationData.Location.Name == "" {
		fmt.Println("The city does not exists")
		return
	}

	var hotMessage string

	if locationData.Current.IsDay == 0 && locationData.Current.TempC > 25 {
		hotMessage += "it's very hot tonight"
	}

	fmt.Println("City name.:", locationData.Location.Name, "\nTemperature.:", locationData.Current.TempC, hotMessage)

}
