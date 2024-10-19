package helpers

var ThresholdTtemperature float64 = 35

func ExceedsThreshold(temp float64) bool {
	threshold := float64(ThresholdTtemperature)
	return CompareTemperatures(temp, threshold)
}

func CompareTemperatures(temp float64, threshold float64) bool {
	return temp > threshold
}
