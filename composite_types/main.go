package main

import "fmt"

func main() {

	// Reverse the whole array a
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println("Array a = [0, 1, 2, 3, 4, 5] reversed will be: ", a)

	// Rotate slice b left by two positions.
	b := []int{0, 1, 2, 3, 4, 5}
	/*
		reverse(b[:2])
		reverse(b[2:])
		reverse(b)
		fmt.Println(b)
	*/

	// slice c
	c := []int{0, 1, 2, 3, 4, 5}

	// slice d
	d := []int{2, 3, 4, 5, 6, 7}

	// Compare
	fmt.Printf(" b == c // %t\n, b == d // %t\n, c == d // %t\n ", equal(b, c), equal(b, d), equal(c, d))

}

//reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i = ", i, "j = ", j, " and s = ", s)
		s[i], s[j] = s[j], s[i]
		fmt.Println("s[i] = ", s[i], "s[j] = ", s[j], " and s = ", s)

	}

}

// Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices ////// contain the same elements.
// We must do comparisons ourselves - Example below
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
