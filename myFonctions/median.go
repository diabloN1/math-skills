package myfonctions

import "sort"

func Median(nums []float64) float64 {
	var sum float64

	// Sort the slice
	sort.Float64s(nums)

	// Check lenght is an even number (2, 4, 6....).
	if len(nums)%2 == 0 {
		sum = nums[(len(nums)/2)] + nums[(len(nums)/2)-1]
		return sum / 2
	}

	return nums[len(nums)/2]
}
