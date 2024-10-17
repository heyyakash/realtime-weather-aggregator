package modals

import "time"

type DailyStats struct {
	AvgTemp         float64   `json:"avgTemp"`
	MaxTemp         float64   `json:"maxTemp"`
	MinTemp         float64   `json:"minTemp"`
	DominantWeather string    `json:"dominantWeather"`
	Day             time.Time `json:"day"`
	City            string    `json:"city"`
}
