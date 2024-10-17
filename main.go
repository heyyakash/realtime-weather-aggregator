package main

import (
	"fmt"
	"log"
	"time"

	"github.com/heyyakash/realtime-weather-aggregator/helpers"
	"github.com/heyyakash/realtime-weather-aggregator/modals"
)

var Cities = []string{"Bangalore", "Delhi"}

var WeatherEventChannel chan modals.WeatherEvent

func Fetch() {
	key := helpers.GetEnv("API_KEY")
	for {
		for _, v := range Cities {
			go func(city string) {
				if err := helpers.FetchWeatherData(city, key, WeatherEventChannel); err != nil {
					log.Printf("Error fetching data for %s : %v", city, err)
				}
			}(v)
		}
		time.Sleep(1 * time.Minute)
	}
}

func main() {
	WeatherEventChannel = make(chan modals.WeatherEvent)

	// go routine for starting server
	go helpers.StartServer()

	// goroutine to start fetching data until stopped
	go Fetch()

	//collect data into channel
	for {
		select {
		case event := <-WeatherEventChannel:
			fmt.Print(event)
		}
	}

}
