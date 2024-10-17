package modals

type WeatherEvent struct {
	EventType   string  `json:"eventType"`
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feelslike"`
	City        string  `json:"city"`
	Dt          int64   `json:"dt"`
	Description string  `json:"description"`
}

func (w *WeatherEvent) ConvertToCelsius() {
	w.Temperature = w.Temperature - 273.16
	w.FeelsLike = w.FeelsLike - 273.16
}

type AlertEvent struct {
	EventType   string  `json:"eventType"`
	Temperature float64 `json:"temperature"`
	City        string  `json:"city"`
}
