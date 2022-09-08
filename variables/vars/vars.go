package vars

import "fmt"

// package-level entity - is visible not only throughout the source file that contains
// its declaration, but throughout all the files of the package

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	BoilingF              = 212.0
	FreezingF             = 32.0
)

// boilingPoint prints the boiling point of water.
func BoilingPoint() {
	var f = BoilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("Boiling Point = %gF or %gC\n", f, c)

}

// Ftoc prints two Fahrenheit-to-Celsius conversions.
func FTOC(f float64) float64 {
	return (f - 32) * 5 / 9
}

//Pointers
/***
A pointer value is the address of a variable. A pointer is thus the location at which a value is
stored. Not every value has an address, but every variable does. With a pointer, we can read
or update the value of a variable in directly, without using or even knowing the name of the
variable, if indeed it has a name.
If a variable is declare d var x int , the expression &x (‘‘address of x ’’) yields a pointer to an
integer variable, that is, a value of type *int , which is pronounced ‘‘pointer to int. ’’ If this
value is called p , we say ‘‘ p points to x , ’’ or equivalently ‘‘ p contains the address of x . ’’ The variable to which p points is written *p . The expression *p yields the value of that
variable, an int , but since *p denotes a variable, it may also appear on the left-hand side of an assig nment, in which case the assig nment updates the var iable.
***/

func TestPointers() {
	x := 1 //Varaible x containing value of 1

	//NB: &x = address of x, yields a pointer to an int in our case i.e yield a value of type *int
	//pronounce "pointer to an int"

	var p = &x // &x = address of x. p is a pointer value and contains the address of variable x

	fmt.Println(*p) // *p denotes a variable and yield the value of that variable in our case the value of x = 1 simply put - *p is an alias of x
	fmt.Println(*p == x)
	*p = 2 // value of x is now 2 i.e x = 2
	fmt.Println(*p == x)
	fmt.Println(x)

	y := &x        // &x = address of x. y is a pointer value and contains the address of variable x
	fmt.Println(y) //yields address of x

	// two pointers are equal if and only if they point to the same variable or both are nil.
	fmt.Println(*y == *p)

	//The new(T) Function
	/***
	The expression new(T) creates an unname d variable of type T , initializes it to the zero value of T , and returns its address, which is a value of type *T .
	***/
	v := new(int) // v is a pointer value of type *int ie v points to an unnamed variable of type int

	fmt.Println(*v) // yields "0"

	*v = 2          // sets the unnamed int variable to 2
	fmt.Println(*v) // yields 2

	q := x //

	fmt.Println(q)
	*v = 9
	q = *v

	fmt.Println(q)
	fmt.Println(q == *v)

}

//type declarations
/***
A type declaration defines a new named type that has the same underlying type as an existing
type. The named type provides a way to separate different and perhaps incompatible uses of
the underlying type so that they can’t be mixed unintentionally.
   *** type name existing-underlying-type ****

For every type T , there is a corresponding conversion operation T(x) that converts the value x
to type T . A conversion from one type to another is allowed if both have the same underlying
type, or if both are unnamed pointer types that point to variables of the same underlying type;
these conversions change the type but not the representation of the value.
***/
//Example - Temperature computations

// defines two typ es, Celsius and Fahrenheit
// Even though both have the same underlying typ e, float64 , they are not the same type, so they
// cannot be compared or combined in arithmetic expressions
type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
