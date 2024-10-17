package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/heyyakash/realtime-weather-aggregator/channels"
	"github.com/heyyakash/realtime-weather-aggregator/modals"
)

var Cities = []string{"Bangalore", "Delhi"}

func Fetch(ctx context.Context) {
	key := GetEnv("API_KEY")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, v := range Cities {
				go func(city string) {
					log.Println("Fetching for : ", city)
					if err := FetchWeatherData(city, key); err != nil {
						log.Printf("Error fetching data for %s : %v", city, err)
					}
				}(v)
			}
			time.Sleep(1 * time.Minute)
		}

	}
}

func FetchWeatherData(city string, key string) error {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, key)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var body modals.WeatherData
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}
	r := modals.WeatherEvent{
		EventType:   "weather_data",
		Temperature: body.Main.Temp,
		FeelsLike:   body.Main.FeelsLike,
		City:        city,
		Dt:          body.Dt,
	}
	r.ConvertToCelsius()
	channels.SSE <- r
	return nil
}
