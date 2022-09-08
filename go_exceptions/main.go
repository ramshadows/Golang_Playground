package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Recall os.Args
	/*
		It’s a slice of string values
		containing all the arguments used when the program is
		run.

		We get the name of the ﬁle to open from the ﬁrst
		command-line argument by accessing os.Args[1] .
		(Remember, the os.Args[0] element is the name
		of the program being run; the actual program arguments
		appear in os.Args[1] and later elements.)
	*/

	priceList, err := GetPrices(os.Args[1])

	// If error encountered, log it and exit

	if err != nil {
		log.Fatal(err)
	}

	// Define a variable to store the sums
	var sum float64 = 0

	// Loop hro the priceList and sum the prices
	for _, price := range priceList {

		sum += price

	}

	fmt.Printf("Total Sum: %.2f\n", sum)
}

// Function OpenFile opens a file and returns a pointer to it
// along with any error encountered

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening ", filename)

	return os.Open(filename)
}

// Function CloseFile closes an opened file

func CloseFile(file *os.File) {
	fmt.Println("Closing file")

	file.Close()
}

//GetPrices reads prices from a text file and returns a slice of the prices as float64
func GetPrices(filename string) ([]float64, error) {
	// stores the prices in a slice
	var prices []float64

	//open the provided file
	file, err := OpenFile(filename)

	// If error was encountered while opening the file return it
	if err != nil {
		log.Fatal(err)
	}

	// The “defer” keyword ensures a function call takes place,
	// even if the calling function exits early.
	// Now, even if GetPrices is given a ﬁle with bad data, it will still
	// close the ﬁle before exiting!

	defer CloseFile(file) // Will run after GetPrices() function exits

	// Declare a new scanner to read the file
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {

		price, err := strconv.ParseFloat(fileScanner.Text(), 64)

		// If error encountered while converting string to float, return it
		// Note: Even if an error gets returned here, CloseFile() will still be called
		if err != nil {
			return nil, err
		}

		prices = append(prices, price)

	}

	// if error encountered scanning the file, return it
	// Note: CloseFile() will be called even if an error is encountered here too
	if fileScanner.Err() != nil {
		return nil, fileScanner.Err()
	}

	// if we got this far, there were no error - return the slice of prices and a nil error
	// Note: And of course CloseFile() is called even if GetPrices() exits without errors
	return prices, nil

}
