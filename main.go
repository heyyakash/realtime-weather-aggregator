package main

import (
	"fmt"
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
			go helpers.FetchWeatherData(v, key, WeatherEventChannel)
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
