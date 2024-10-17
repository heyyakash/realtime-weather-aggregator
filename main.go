package main

import (
	"log"
	"net/http"

	"github.com/heyyakash/realtime-weather-aggregator/configs"
	"github.com/heyyakash/realtime-weather-aggregator/handlers"
)

func StartServer() {
	http.HandleFunc("/events", handlers.EventsHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting the server : ", err)
	}
}

func main() {
	// connected to the database
	configs.ConnectDB()

	//start the http server
	StartServer()
}
