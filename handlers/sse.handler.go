package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	log.Print("Client connected")
	ctx := r.Context()

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			log.Print("Client disconnected")
			return
		default:
			fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
			w.(http.Flusher).Flush()
			time.Sleep(2 * time.Second)
		}
	}
}
