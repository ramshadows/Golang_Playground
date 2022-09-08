package stringerInterface

import "fmt"

/*
Stringer Interface from the fmt Package
Many functions in the fmt package check whether the
values passed to them satisfy the Stringer interface, and
call their String methods if so.
This includes the Print , Println , and Printf functions and more.
     ****************
	 type Stringer interface {
		 String() string
	 }
*/

// Define two new types each with an underlying type float64
// Note: A deﬁned type supports all the same operations as its
// underlying type. - In our case float64
type Liters float64
type Gallons float64
type Milliliters float64

/*
Now, any time we pass Gallons , Liters , and Milliliters
values to Println (or most other fmt functions), their String
methods will be called, and the return values used in the
output. We’ve set up a useful default format for printing
each of these types!
*/

// Make Liter satisfy Stringer

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

// Make Gallons satisfy Stringer

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

// Make Milliliters satisfy Stringer

func (ml Milliliters) String() string {
	return fmt.Sprintf("%0.2f ML", ml)
}
