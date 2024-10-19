package handlers

import (
	"net/http"
	"strconv"

	"github.com/heyyakash/realtime-weather-aggregator/helpers"
)

func UpdateThreshold(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		newValue := r.URL.Query().Get("threshold")
		newThreshold, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			http.Error(w, "Send a valid threshold value", http.StatusBadRequest)
			return
		}
		helpers.ThresholdTtemperature = newThreshold
		w.Write([]byte("success"))
	}
}
