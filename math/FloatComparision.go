package Math

import (
	"math"
)

const Threshold = 0.00001

func CompareFloatsWithThreshold(a, b, t float64) bool {
	return math.Abs(a-b) <= t
}

func CompareFloats(a, b float64) bool {
	return CompareFloatsWithThreshold(a, b, Threshold)
}
