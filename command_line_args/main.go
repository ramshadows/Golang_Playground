package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]

	var sum float64 = 0

	for _, argument := range arguments {
		nums, err := strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(err)
		}

		sum += nums
	}

	sampleCount := float64(len(arguments))

	fmt.Printf("Average of arguments is: %.2f\n", sum/sampleCount)
}
