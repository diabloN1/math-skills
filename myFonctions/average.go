package myfonctions

func Average(nums []float64) float64 {
	var sum, average float64

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	average = sum / float64(len(nums))

	return average
}
