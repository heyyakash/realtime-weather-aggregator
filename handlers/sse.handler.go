package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/heyyakash/realtime-weather-aggregator/channels"
	"github.com/heyyakash/realtime-weather-aggregator/configs"
	"github.com/heyyakash/realtime-weather-aggregator/helpers"
	"github.com/heyyakash/realtime-weather-aggregator/modals"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx, cancel := context.WithCancel((r.Context()))
	defer cancel()
	go helpers.Fetch(ctx)

	for {
		select {
		case <-ctx.Done():
			log.Print("Client disconnected")
			return
		case msg := <-channels.SSE:

			// insert data into mongodb if only it is weather data and not alert
			go func(msg interface{}) {
				data, ok := msg.(modals.WeatherEvent)
				if ok {
					if data.EventType == "weather_data" {
						configs.WeatherDataCollection.InsertOne(context.TODO(), data)
					}
				}
			}(msg)

			res, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error marshalling message: %v", err)
				continue
			}
			fmt.Fprintf(w, "data: %s\n\n", res)
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		}
	}
}
