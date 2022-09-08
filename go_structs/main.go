package main

import (
	"fmt"
	"git_repo/go_playground/playground/go_structs/exercise1"
)

func main() {
	var myCar exercise1.Car
	exercise1.Subscribe()
	exercise1.ShowInfo(myCar)

	bolt := exercise1.MinimumOrder("Hex Bolts")

	fmt.Printf("Part Desc: %s, Part Cost: %.2f\n", bolt.Description, bolt.Cost)

	// This is no longer a struct - it is a pointer to a struct
	subscriber1 := exercise1.DefaultSubscriber("Charles Kith")

	exercise1.ApplyDiscount(subscriber1) // This removes the "Address of operator - (&)" and updates the rate
	exercise1.PrintInfo(subscriber1)

	subscriber2 := exercise1.DefaultSubscriber("Aaron Hickey")
	exercise1.PrintInfo(subscriber2)

	// Populating the HomeAdress field
	address1 := exercise1.Address{
		City:       "Nairobi",
		County:     "Kiambu",
		PostalCode: "001000-6000",
	}

	var emp1 exercise1.Employee
	emp1.HomeAddress = address1

	fmt.Println(emp1.HomeAddress)
	fmt.Println(emp1.HomeAddress.PostalCode)

	// Accessing the Anonymous struct in Employee2
	var emp2 exercise1.Employee2
	emp2.Address = address1
	fmt.Println(emp2.Address)
	fmt.Println("County: ", emp2.Address.County)

	// Alternative
	fmt.Println(emp2.City)
	fmt.Println("County: ", emp2.County)

}
