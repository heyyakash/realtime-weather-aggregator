package modals

type WeatherEvent struct {
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feelslike"`
	City        string  `json:"city"`
	Dt          int64   `json:"dt"`
}

func (w *WeatherEvent) ConvertToCelsius() {
	w.Temperature = w.Temperature - 273.16
	w.FeelsLike = w.FeelsLike - 273.16
}
