package helpers_test

import (
	"fmt"
	"testing"

	"github.com/heyyakash/realtime-weather-aggregator/helpers"
	"github.com/heyyakash/realtime-weather-aggregator/modals"
)

func TestFetchWeatherData(t *testing.T) {
	api_key := helpers.GetEnv("API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", "Delhi", api_key)
	var body modals.WeatherData
	if err := helpers.FetchWeatherGetRequest(&body, url); err != nil {
		t.Error(err)
	}
}
