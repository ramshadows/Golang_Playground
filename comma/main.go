// comma inserts commas in a non-negative decimal integer string.

package main

import (
	"bytes"
	"fmt"
)

func main() {
	str_num := comma("11870500")

	fmt.Printf("Add comma to 11870500 becomes: %s\n", str_num)

	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

}

/*
Takes a string representation of an integer, such as "12345" , and insert commas every three places, as in "12,345"
*/
// The argument to comma is a str ing .
func comma(s string) string {
	//If its length is less than or equal to 3, no comma is necessary
	n := len(s)

	if n <= 3 {
		return s
	}
	// Other wise, comma calls itself recursively with a substring consisting of all but the last
	// three characters, and appends a comma and the last three characters to the result of the
	//recursive call.

	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprintf(values) but adds commas.
// Buffer starts out empty but grows as data of types like string , byte , and []byte
// are written to it.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
