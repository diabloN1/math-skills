package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"mathskills/myFonctions"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Command not valid :\nTry Instead \"go run your-program.go data.txt\"")
	}
	file := myfonctions.Read(os.Args[1])
	data := strings.Split(file[:len(file)-1], "\n")
	floatedData := myfonctions.StringsToFloats(data)
	fmt.Printf("Average: %.0f\n", math.Round(myfonctions.Average(floatedData))) //f stands for "floating point."
	fmt.Printf("Median: %.0f\n", math.Round(myfonctions.Median(floatedData)))
	fmt.Printf("Variance: %.0f\n", math.Round(myfonctions.Variance(floatedData)))
	fmt.Printf("StandardDeviation: %.0f\n", math.Round(myfonctions.StandardDeviation(floatedData)))
}
