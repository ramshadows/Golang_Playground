package main

import "fmt"

// Define two new types each with an underlying type float64
// Note: A deﬁned type supports all the same operations as its
// underlying type. - In our case float64
type Liters float64
type Gallons float64
type Milliliters float64

func main() {
	// Define a variable of type Gallon
	var carFuel Gallons

	// Define a variable of type Liters
	var truckFuel Liters

	// Convert a float to Gallon
	carFuel = Gallons(10.0)

	// Convert a float to Liters
	truckFuel = Liters(240.0)

	fmt.Println("Car Fuel: ", carFuel, " Gallons.  Truck Fuel: ", truckFuel, " Liters")

	// Convert Liters to Gallons
	carFuel = Gallons(Liters(40.0) * 0.264)

	// Convert Gallons to Liters
	truckFuel = Liters(Gallons(63.0) * 3.785)

	fmt.Printf("Car Fuel: %.2f Gallons.  Truck Fuel: %.2f Liters\n", carFuel, truckFuel)

	// A deﬁned type can be used in operations together with literal values:
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.5)

	// But deﬁned types cannot be used in operations together
	// with values of a diﬀerent type, even if the other type has
	// the same underlying type.
	// If you want to add a value in Liters to a value in Gallons ,
	// you’ll need to convert one type to match the other ﬁrst.

	//fmt.Println(Liters(1.2) + Gallons(5.9)) -> Invalid Operation

	//Converting between types using functions
	fuel := 10.0
	fmt.Printf("%.1f Liters equals %.1f Gallons\n", fuel, toGallons(Liters(fuel)))
	fmt.Printf("%.1f Gallons equals %.1f Liters\n", fuel, toLiters(Gallons(fuel)))

	// Using Defined Methods
	soda := Liters(2.0)
	fmt.Printf("%.1f Liters equals %.1f Gallons\n", soda, soda.toGallons())

	// Create a Milliliters value
	water := Milliliters(6.0)

	// Convert Milliliters to Gallons
	fmt.Printf("%.1f Milliliters equals %.1f Gallons\n", water, water.toGallons())

	// Create a Gallon value
	milk := Gallons(2)

	fmt.Printf("%.1f Gallons equals %.1f Milliliters\n", milk, milk.toMilliliters())

	fmt.Printf("%.1f Gallons equals %.1f Liters\n", milk, milk.toLiters())

}

// Converting between types using functions

func toGallons(a Liters) Gallons {
	return Gallons(a * 0.264)
}

func toLiters(a Gallons) Liters {
	return Liters(a * 3.785)
}

// Using Defined Methods
// Notice names can be identical if they are on different types
func (m Milliliters) toGallons() Gallons {
	return Gallons(m * 0.264)

}

func (a Liters) toGallons() Gallons {
	return Gallons(a * 0.264)

}

func (g Gallons) toMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)

}
