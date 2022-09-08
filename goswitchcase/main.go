package main

import (
	"fmt"
	goswitchcase "git_repo/go_playground/playground/goswitchcase/switch_case"
	"math/rand"
	"regexp"
	"time"
)

func main() {
	goswitchcase.GoSwitchCase()

	region, continent := goswitchcase.Location("Kigali")
	fmt.Printf("John works in %s, %s\n", region, continent)

	// A switch can also invoke a function. From that function, you can write
	// case statements for possible return values. For example, the following
	// code calls the time.Now() function. The output that it prints depends on the
	// current weekday.
	// When you call a function from a switch statement, you can modify its logic
	// without changing the expression because you always validate what the function returns.
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())

	// Also, you can call a function from a case statement.
	// Notice that the switch block has no validating expression

	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	// you can omit a condition in a switch statement like you do in an if statement.
	// This pattern is like comparing a true value as if you were forcing the switch
	// statement to run all the time.
	// he program always runs this type of switch statement because the condition is always true. 
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}
}
