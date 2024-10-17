package main

import (
	"fmt"

	"github.com/heyyakash/realtime-weather-aggregator/helpers"
)

var Cities = []string{"Bangalore", "Delhi"}

func FetchWeatherData() {
	key := helpers.GetEnv("API_KEY")
	body := helpers.FetchWeatherData("Bangalore", key)
	fmt.Print(body.Weather[0].Description)
}

func main() {
	FetchWeatherData()
}
