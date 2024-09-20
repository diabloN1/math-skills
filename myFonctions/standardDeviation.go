package myfonctions

import "math"

func StandardDeviation(nums []float64) float64 {
	standardDeviation := math.Sqrt(Variance(nums))

	return standardDeviation
}
