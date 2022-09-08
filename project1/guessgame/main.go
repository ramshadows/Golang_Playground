package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get the current time and date as an integer
	// The rand.Seed function expects an integer, so we can’t pass
	// it a Time value directly. Instead, we call the Unix method on
	// the Time , which will convert it to an integer.
	seconds := time.Now().Unix()

	// Use the second as seed value to enable generate unique rand each time
	rand.Seed(seconds)

	// Call the rand.Intn to generate random number 1 to 100
	// Note: generates range 0-99 thus add 1 to get range 0-100
	target := rand.Intn(100) + 1

	fmt.Println("I've choosen a random integer.")
	fmt.Println("Can you guess it?")

	fmt.Printf("Number to guess is: %d\n", target)

	//Creates a buffer io Reader which let us read keybaord input
	reader := bufio.NewReader(os.Stdin)

	success := false
	trials := 3

	// Loop to prompt user for ten guessess
	for guesses := 0; guesses < trials; guesses++ {
		fmt.Println("You have ", trials-guesses, " guesses left")

		//Prompts a user enter guess number from keyboard
		fmt.Print("Make a guess: ")

		// Read what the user types up untill they press enter key
		input, err := reader.ReadString('\n')

		// if encounter an error , read the error message and exit code
		if err != nil {
			log.Fatal(err)
		}

		// Trim new character, white spaces both right and left from the input
		input = strings.TrimSpace(input)

		// Convert the input from a string to an Integer
		guessedNo, err := strconv.Atoi(input)

		// if encounter an error , read the error message and exit code
		if err != nil {
			log.Fatal(err)
		}

		if guessedNo < target {
			fmt.Println("Oops. Your guess was: LOW")

		} else if guessedNo > target {
			fmt.Println("Oops., Your guess was: HIGH")

		} else {
			success = true
			fmt.Println("Good Job, Your guessed it: RIGHT")
			break

		}

	}

	if !success {
		fmt.Printf("Sorry, You've exhausted your allowed guesses.\n")
		fmt.Println("You didn't guess my number. Target guess was: ", target)
	}

}
