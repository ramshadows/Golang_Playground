package main

import (
	"fmt"
	"git_repo/go_playground/playground/go_setter_methods/calender"

	"log"
)

func main() {
	// Create struct variable of type Date
	date := calender.Date{}

	yearErr := date.SetYear(2022)

	if yearErr != nil {
		log.Fatal(yearErr)
	}

	// Set the Year, Month and Day

	monthErr := date.SetMonth(5)

	if monthErr != nil {
		log.Fatal(monthErr)
	}

	dayErr := date.SetDay(17)

	if dayErr != nil {
		log.Fatal(dayErr)
	}

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	// Create an Event struct variable
	event := calender.Event{}

	titleErr := event.SetTitle("Birthday Party")

	if titleErr != nil {
		log.Fatal(titleErr)
	}

	// You can use the event variable to access exported methods of the embedded type
	fmt.Println(event.Title())
	fmt.Println(event.Date.Year()) // Using the chained dot operator
	fmt.Println(event.Month())     // Accessing directly
	fmt.Println(event.Day())

}
