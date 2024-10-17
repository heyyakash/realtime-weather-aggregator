package helpers

import (
	"testing"
)

type Test struct {
	CurrentTemp float64
	Threshold   float64
	Expect      bool
}

func TestCompareTemperatures(t *testing.T) {
	tests := []Test{
		{CurrentTemp: 24, Threshold: 23, Expect: true},
		{CurrentTemp: 24, Threshold: 24, Expect: false},
		{CurrentTemp: 24, Threshold: 25, Expect: false},
	}

	for _, test := range tests {
		res := CompareTemperatures(test.CurrentTemp, test.Threshold)
		if res != test.Expect {
			t.Error("For Current Temp : ", test.CurrentTemp, " and Threshold Temp : ", test.Threshold, " Expected  : ", test.Expect, "Got : ", res)
		}
	}
}
