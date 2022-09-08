package main

import (
	"bufio"
	"fmt"
	"git_repo/go_playground/playground/read_text_files/readTextFile"

	"log"
	"os"
	"strconv"
)

var filePath = "/home/ram/Golang_Playground/src/git_repo/go_playground/chapter2/read_text_files/"
var data = filePath + "prices.txt"

func main() {

	prices, err := readTextFile.GetPrices(data)

	// If error encounterd reading the file - return it and exit
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prices)

	// sum the prices
	var sum float64 = 0.00

	for _, price := range prices {
		sum += price

	}

	sampleCount := float64(len(prices))

	fmt.Printf("Average price is: %.2f\n", sum/sampleCount)

	// Pass data.txt as a command line argument
	arguments := os.Args[1:]

	var sum2 float64 = 0

	for _, argument := range arguments {
		fmt.Println(argument)

		//open the provided file
		file, err := os.Open(argument)

		// If error was encountered while opening the file return it
		if err != nil {
			log.Fatal(err)
		}

		// Declare a new scanner to read the file
		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {

			price2, err := strconv.ParseFloat(fileScanner.Text(), 64)

			if err != nil {
				log.Fatal(err)
			}

			sum2 += price2

		}

	}

	sampleCount2 := float64(len(arguments))

	fmt.Printf("Average prices2 is: %.2f\n", sum2/sampleCount2)

}
