package myfonctions

import "math"

func Variance(nums []float64) float64 {
	var sum, variance float64

	// Variance is the sum of each number - average power 2
	for i := 0; i < len(nums); i++ {
		sum += math.Pow(nums[i]-Average(nums), 2)
	}
	variance = sum / float64(len(nums))

	return variance
}
