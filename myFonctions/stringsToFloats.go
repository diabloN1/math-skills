package myfonctions

import (
	"log"
	"strconv"
)

func StringsToFloats(data []string) []float64 {
	floats := []float64{}

	for i := range data {
		float, err := strconv.ParseFloat(data[i], 64)
		if err != nil {
			log.Fatalln("There was an error while float parsing :", data[i])
		}
		floats = append(floats, float)
	}

	return floats
}
