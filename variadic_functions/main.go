/*
A variadic function is one that can be called with a varying number of arguments. To make a
function variadic, use an ellipsis ( ... ) before the type of the last (or only) function parameter in the function declaration.

The last parameter of a variadic function receives the
variadic arguments as a slice, which the function can then
process like any other slice.

A function can take one or more nonvariadic arguments as
well. Although a function caller can omit variadic
arguments (resulting in an empty slice), nonvariadic
arguments are always required; it’s a compile error to omit
those.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Maximum of [30, 40, 60, 10, 4] is: ", maximum(30, 40, 60, 10, 4))

	fmt.Println(inRange(1, 10, 4.5, 6, 2, 6, 12, 10, 0))
	fmt.Println(inRange(-10, 10, 4.5, 6, -2, 6, 12, -10, 0))

	//Passing slices to variadic functions
	/*
			When calling a variadic function, simply add an ellipsis
		( ... ) following the slice you want to use in place of variadic
		arguments.
	*/
	mySlice := []float64{-10, 10, 4.5, 6, -2, 6, 12, -10, 0}

	fmt.Println(inRange(1, 10, mySlice...))

}

func maximum(nums ...float64) float64 {
	// Start with a very low number - Negative Infinity
	max := math.Inf(-1)

	for _, num := range nums {

		if num > max {
			max = num

		}

	}

	return max
}

/*
Here’s an inRange function that takes a minimum value, a
maximum value, and any number of additional float64
arguments. It will discard any argument that is below the
given minimum or above the given maximum, returning a
slice containing only the arguments that were in the
speciﬁed range.
*/

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64

	for _, num := range numbers {
		if num >= min && num <= max {
			result = append(result, num)
		}

	}

	return result
}
