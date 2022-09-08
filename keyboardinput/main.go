package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Prompts a user to enter a grade from keyboard
	fmt.Print("Enter Student grade then press enter: ")

	//Creates a buffer io Reader which let us read keybaord input
	reader := bufio.NewReader(os.Stdin)

	// Read what the user types up untill they press enter key
	input, err := reader.ReadString('\n')

	// if encounter an error , read the error message and exit code
	if err != nil {
		log.Fatal(err)
	}

	// Trim new character, white spaces both right and left from the input
	input = strings.TrimSpace(input)

	// Convert the input from a string to a float64
	grade, err := strconv.ParseFloat(input, 64)

	// if encounter an error , read the error message and exit code
	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 80 {
		status = "Pass"
	} else {
		status = "Fail"
	}

	fmt.Println("A grade of: ", grade, "is: ", status)

}
