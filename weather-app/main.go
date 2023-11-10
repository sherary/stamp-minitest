package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	API_KEY        = ""
	geoAPIEndpoint = "http://api.openweathermap.org/geo/1.0/direct"
	weatherAPI     = "http://api.openweathermap.org/data/2.5/forecast"
)

type GeoResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ForecastResponse struct {
	Cod  string     `json:"cod"`
	Cnt  int        `json:"cnt"`
	List []Forecast `json:"list"`
}

type Forecast struct {
	Dt   int64        `json:"dt"`
	Main TempInMetric `json:"main"`
}

type TempInMetric struct {
	Temp    float32 `json:"temp"`
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
}

var httpClient = &http.Client{}

func fetchGeoData(city string) ([]GeoResponse, error) {
	url := fmt.Sprintf("%s?q=%s&limit=5&appid=%s", geoAPIEndpoint, city, API_KEY)
	var geoResponse []GeoResponse
	err := fetchData(url, &geoResponse)
	return geoResponse, err
}

func fetchWeatherData(lat, lon float64) (ForecastResponse, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=metric", weatherAPI, lat, lon, API_KEY)
	var weatherResponse ForecastResponse
	err := fetchData(url, &weatherResponse)
	return weatherResponse, err
}

func fetchData(url string, target interface{}) error {
	response, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

func main() {
	city := "Jakarta"
	geoData, err := fetchGeoData(city)
	if err != nil {
		fmt.Println("Error fetching geo data:", err)
		return
	}

	if len(geoData) > 0 {
		geo := geoData[0]
		lat, lon := geo.Lat, geo.Lon

		weatherData, err := fetchWeatherData(lat, lon)
		if err != nil {
			fmt.Println("Error fetching weather data:", err)
			return
		}

		for i, forecast := range weatherData.List {
			if i%8 == 0 {
				timestamp := time.Unix(forecast.Dt, 0)
				formattedDate := timestamp.Format("Mon, 02 Jan 2006:")
				fmt.Printf("%s %.2f %s\n", formattedDate, forecast.Main.Temp, "°C")
			}
		}
	}
}
