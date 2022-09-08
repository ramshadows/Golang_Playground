package main

import (
	"fmt"
	"log"
)

func main() {

	fridge := refrigerator{"Milk", "Pizza", "Sausages"}

	for _, foodItem := range []string{"Milk", "Bananas"} {

		err := fridge.findFood(foodItem)

		if err != nil {
			log.Fatal(err)
		}

	}

}

// Find() function searches the underlying slice for a particular food
// Returns true if found otherwise false

func find(foodType string, foodList []string) bool {
	for _, food := range foodList {
		if foodType == food {
			return true

		}

	}

	return false

}

// The refrigerator type has an underlying type of slice of strings which will
// hold the names of the foods the refrigerator contains

type refrigerator []string

// refrigerator type has a method open() that simulates refrigerator door opening

func (r refrigerator) open(item string) {
	fmt.Println("Opening the Refrigerator to search for ", item)
}

// refrigerator type has a method close() that simulates refrigerator door closing

func (r refrigerator) close() {
	fmt.Println("Closing the Refrigerator door")
}

/*
// The FindFood method calls Open to open the door,
// calls a find function we’ve written to search the underlying
// slice for a particular food, and then calls Close to close the
// door again.
*/

func (r refrigerator) findFood(food string) error {
	r.open(food)

	// Now, even if findF() does not find the item we are looking for,
	// the Refrigerator door will be closed after the food is not found.
	// Note: You can only defer function and method calls
	defer r.close()

	if find(food, r) {
		fmt.Println("Found: ", food)

	} else {
		fmt.Printf("%s not found\n", food)
	}

	// r.close() would still be called even if findFood() exited successifully ie
	// the food was found. still close the Refrigerator door
	return nil
}
