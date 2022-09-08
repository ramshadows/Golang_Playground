package errorInterface

import "fmt"

/*
The error type is just an interface! It looks
something like this:

type error interface {
	Error() string
}

That means you can deﬁne your own types and use them anywhere an error
value is required!

*/

// Defines a type with an underlying type of float64

type OverheatingError float64

func (heating OverheatingError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", heating)

}

// checkTemperature function that uses OverheatError
// Note: We are speciffying that the function returns an ordinary error - error
/*
The error type is a “predeclared identiﬁer,” like int or
string . And so, like other predeclared identiﬁers, it’s not
part of any package. It’s part of the “universe block,”
meaning it’s available everywhere, regardless of what
package you’re in.
*/

func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe

	// If the actual temperature is in excesses of the safe temp
	if excess > 0 {
		// return an overheating error that records the excess
		return OverheatingError(excess)
	}

	// Otherwise return nil meaning no overheating
	return nil

}
