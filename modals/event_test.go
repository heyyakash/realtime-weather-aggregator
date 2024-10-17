package modals

import (
	"testing"
)

func TestConvertFromKelvinToCentigrade(t *testing.T) {
	stat := WeatherEvent{
		Temperature: 273.16,
		City:        "Bangalore",
		FeelsLike:   273.16,
		Dt:          1729178565,
		EventType:   "weather_data",
		Description: "Hazy",
	}

	stat.ConvertToCelsius()
	temp := stat.Temperature
	fl := stat.FeelsLike

	if temp != float64(0) {
		t.Errorf("Expected temperature to be 0 but got : %f", temp)
	}
	if fl != float64(0) {
		t.Errorf("Expected feels like to be 0 but got : %f", temp)
	}
}
