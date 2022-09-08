package main

import (
	"fmt"
	"git_repo/go_playground/playground/go_buitIn_interfaces/errorInterface"
	"git_repo/go_playground/playground/go_buitIn_interfaces/stringerInterface"

	"log"
)

func main() {
	var err error = errorInterface.CheckTemperature(29.36, 100.00)

	if err != nil {
		log.Fatal(err)
	}

	// Output -> Overheating by 29.36 degrees

	fmt.Println(stringerInterface.Liters(40.7896))      // Output -> 40.79 L
	fmt.Println(stringerInterface.Gallons(40.7896))     // Output -> 40.79 gal
	fmt.Println(stringerInterface.Milliliters(40.7896)) // Output -> 40.79 ML
}
