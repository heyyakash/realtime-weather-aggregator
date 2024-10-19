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

var Cities = []string{"Delhi", "Mumbai", "Chennai", "Bangalore", "Kolkata", "Hyderabad"}

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
	var body modals.WeatherData
	if err := FetchWeatherGetRequest(&body, url); err != nil {
		log.Print(err)
	}
	r := modals.WeatherEvent{
		EventType:   "weather_data",
		Temperature: body.Main.Temp,
		FeelsLike:   body.Main.FeelsLike,
		City:        city,
		Dt:          body.Dt,
		Description: body.Weather[0].Main,
	}
	r.ConvertToCelsius()
	channels.SSE <- r

	if ExceedsThreshold(r.Temperature) {
		alert := modals.AlertEvent{
			EventType:   "alert_data",
			Temperature: r.Temperature,
			City:        r.City,
		}
		channels.SSE <- alert
	}

	return nil
}

func FetchWeatherGetRequest(body *modals.WeatherData, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}
	return nil
}
