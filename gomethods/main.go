package main

import (
	"fmt"
	"strings"
)

/*
A method in Go is a special type of function with a simple difference:
you have to include an extra parameter before the function name.
This additional parameter is known as the receiver.

Methods are useful when you want to group functions and tie them to a custom type.

**** Declare methods ****
func (variable type) MethodName(parameters ...) returnType-Optional{
    // method functionality
}

However, before you can declare a method, you have to create a struct, or
any other underlying type such as string, float64, int64 e.t.c
*/

type Triangle struct {
	size int
}

type Square struct {
	size int
}

// // Declare methods for other types
type UpperCase string

// Embed struct
type coloredTriangle struct {
    Triangle  // embede struct
    color string
}


// perimeter method that calculate the perimeter of triangle
// Notice the receiver type is of type Triangle struct above
func (t Triangle) perimeter() int {
	return t.size * 3
}

func (s Square) perimeter() int {
	return s.size * 4
}

// use pointers update the receiver variable in a method.
func (t *Triangle) doubleSize() {
	t.size *= 2
}

func (myString UpperCase) toUpperCase() string {
	return strings.ToUpper(string(myString))
}

// wrapper method on colouredTrinagle to change its own implementation of perimeter method
func (t2 coloredTriangle) perimeter() int {
	return t2.size * 3 * 2 // doubles the perimeter
}



func main() {
	t := Triangle{size: 3}

	// If you try to call the perimeter() function as you ordinarily would, it won't
	// work because the function's signature says that it needs a receiver.
	// That's why the only way to call that method is to declare a struct first,
	// which gives you access to the method. This means that you could even have
	// the same name for a method as long as it belongs to a different struct.

	s := Square{size: 4}

	fmt.Println("Perimeter of trinagle with is: ", t.perimeter())
	fmt.Println("Perimeter of square with is: ", s.perimeter())

	// call doubleSize() to double the size of triangle
	t.doubleSize()

	fmt.Println("Triangle size after doubling: ", t.size)
	fmt.Println("Perimeter of trinagle after doubling the size: ", t.perimeter())
	fmt.Println("Triangle size after doubling: ", t.size)

	myString := UpperCase("Learning Go!")
    fmt.Println(myString)
    fmt.Println(myString.toUpperCase())

	// You can call methods of the embedded struct even if the receiver is different.
	t2 := coloredTriangle{Triangle{3}, "blue"}
    fmt.Println("Size:", t2.size)
    fmt.Println("Perimeter", t2.perimeter())
}
