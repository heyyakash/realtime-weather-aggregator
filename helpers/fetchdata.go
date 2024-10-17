package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heyyakash/realtime-weather-aggregator/modals"
)

func FetchWeatherData(city string, key string) modals.WeatherData {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, key)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var body modals.WeatherData
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(err)
	}
	return body
}
