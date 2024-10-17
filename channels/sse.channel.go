package channels

import "github.com/heyyakash/realtime-weather-aggregator/modals"

var SSE = make(chan interface{})

func SendSSEEvent(event modals.WeatherEvent) {
	SSE <- event
}

func CloseSSEChannel() {
	close(SSE)
}
