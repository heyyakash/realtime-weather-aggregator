package helpers

import (
	"log"
	"strconv"
)

func ExceedsThreshold(temp float64) bool {
	threshold, err := strconv.ParseFloat(GetEnv("THRESHOLD"), 64)
	if err != nil {
		log.Fatal("Invalid Threshold value")
	}
	return temp >= threshold
}
