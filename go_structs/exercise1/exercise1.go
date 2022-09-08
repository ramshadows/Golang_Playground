package exercise1

import "fmt"

/*
#Defined types and structs

Type deﬁnitions allow you to create types of your own.
They let you create a new deﬁned type that’s based on an
underlying type.

To write a type deﬁnition, use the type keyword, followed by
the name for your new deﬁned type, and then the
underlying type you want to base it on. If you’re using a
struct type as your underlying type, you’ll use the struct
keyword followed by a list of ﬁeld deﬁnitions in curly
braces, just as you did when declaring struct variables.

        ***********
		type myType struct {
			// field names
		}
*/
// Example

type Part struct {
	Description string
	Cost        float64
	Quantity    int
}

type Car struct {
	carMake           string
	model             string
	bodyShape         string
	yearOfManufacture int
}

type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func Subscribe() {
	// Declare a struct type
	var sub struct {
		name   string
		rate   float64
		active bool
	}

	// Assign values of appropriate type to each field
	// Notce we assign values using the dot(.) operator
	sub.name = "Peter"
	sub.rate = 400
	sub.active = true

	// Print the struct field values
	// Notce we retrieve values using the dot(.) operator
	fmt.Println("Subscriber name: ", sub.name)
	fmt.Println("Monthly subscription rate: ", sub.rate)
	fmt.Println("Member is active: ", sub.active)

	toyota := Car{
		carMake:           "Cross Road",
		model:             "SUV",
		bodyShape:         "Station Wagon",
		yearOfManufacture: 2020,
	}

	fmt.Println("Car Make: ", toyota.carMake)
	fmt.Println("Car Model: ", toyota.model)
	fmt.Println("Car Body Shape: ", toyota.bodyShape)
	fmt.Println("Car Year of Manufacture: ", toyota.yearOfManufacture)

	bolts := Part{
		Description: "Hex Bolts",
		Cost:        89.0,
		Quantity:    20,
	}

	fmt.Println("Part Description: ", bolts.Description)
	fmt.Println("Part Cost: ", bolts.Cost)
	fmt.Println("Part Quantity: ", bolts.Quantity)

}

// Using deﬁned types with functions
// Defined Types also work for function parameters and return
// values.

func ShowInfo(myCar Car) {
	fmt.Println("Car Make: ", myCar.carMake)
	fmt.Println("Car Model: ", myCar.model)
	fmt.Println("Car Body Shape: ", myCar.bodyShape)
	fmt.Println("Car Year of Manufacture: ", myCar.yearOfManufacture)

}

// Here MinimumOrder as a return type of struct
func MinimumOrder(description string) Part {
	var p Part

	p.Description = description
	p.Cost = 100.0

	return p
}

// Passing struct pointer as function parameters
func PrintInfo(s *Subscriber) {
	fmt.Println("Subscriber name: ", s.name)
	fmt.Println("Monthly subscription rate: ", s.rate)
	fmt.Println("Member is active: ", s.active)

}

// Returning struct pointer
func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber

	s.name = name
	s.rate = 5.99
	s.active = true

	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.rate = 4.99
}

// struct can hold other structs
type Employee struct {
	FullName    string
	Salary      float64
	HomeAddress Address
}

// Using Anonymous struct
type Employee2 struct {
	FullName    string
	Salary      float64
	Address     // Notice we have deleted the field name
}


type Address struct {
	City       string
	County     string
	PostalCode string
}


