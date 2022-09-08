package readTextFile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//GetPrices reads prices from a text file and returns a slice of the prices as float64
func GetPrices(filename string) ([]float64, error) {
	// stores the prices in a slice
	var prices []float64

	//open the provided file
	file, err := os.Open(filename)

	// If error was encountered while opening the file return it
	if err != nil {
		log.Fatal(err)
	}

	// Declare a new scanner to read the file
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {

		price, err := strconv.ParseFloat(fileScanner.Text(), 64)

		// If error encountered while converting string to float, return it
		if err != nil {
			return nil, err
		}

		prices = append(prices, price)

	}

	// close the file after reading
	err = file.Close()

	// if error encounterd closing the file return it
	if err != nil {
		return nil, err
	}

	// if error encountered scanning the file, return it
	if fileScanner.Err() != nil {
		return nil, fileScanner.Err()
	}

	// if we got this far, there were no error - return the slice of prices and a nil error
	return prices, nil

}

// Function GetStrings reads a file containing strings and returns strings and an error if any
func GetStrings(datafile string) ([]string, error) {

	// stores the prices in a slice
	var myStrings []string

	//open the provided file
	file, err := os.Open(datafile)

	// If error was encountered while opening the file return it
	if err != nil {
		log.Fatal(err)
	}

	// Declare a new scanner to read the file
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {

		lines := fileScanner.Text()
		myStrings = append(myStrings, lines)

		// If error encountered while converting string to float, return it
		if err != nil {
			return nil, err
		}

	}

	// close the file after reading
	err = file.Close()

	// if error encounterd closing the file return it
	if err != nil {
		return nil, err
	}

	// if error encountered scanning the file, return it
	if fileScanner.Err() != nil {
		return nil, fileScanner.Err()
	}

	// if we got this far, there were no error - return the slice of strings and a nil error
	return myStrings, nil

}
