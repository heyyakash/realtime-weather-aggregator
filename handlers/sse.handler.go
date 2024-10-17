package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/heyyakash/realtime-weather-aggregator/channels"
	"github.com/heyyakash/realtime-weather-aggregator/helpers"
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
			log.Print(msg)
			res, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error marshalling message: %v", err)
				continue // Skip this iteration on error
			}
			fmt.Fprintf(w, "data: %s\n\n", res) // Ensure data is prefixed correctly for SSE
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		}
	}
}
