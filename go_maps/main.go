package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declares an empty map variable
	var nilMap map[string]int

	fmt.Printf("%#v \n", nilMap) //Output - map[string]int(nil)
	//nilMap["Gold"] = 3 // map is nil can't add values #Output - panic: assignment to entry nil map

	// Creates a map
	nilMap = make(map[string]int)
	nilMap["Bronze"] = 10        //Now you can add values
	fmt.Printf("%#v \n", nilMap) //Output is: nilMap[string]int("Bronze" : 10)

	// Declare and creates a map
	var nilMap2 map[string]int = make(map[string]int)
	nilMap2["Gold"] = 3

	fmt.Printf("%#v \n", nilMap2)

	// Using the statu1 func
	status1("peter") // A map key with value zero assigned - #Output - "peter has failed"
	status1("john")  // A map key with no value assigned assigned - #Output - "john has failed"

	// Using the statu2 func
	status2("peter") // A map key with value zero assigned - #Output - "peter has failed"
	status2("john")  // A map key with no value assigned assigned - #Output - "No grade assigned to john"

	myData := []string{"a", "c", "e", "a", "e"}

	counter := make(map[string]int)

	for _, item := range myData {

		counter[item]++

	}
	fmt.Println(counter)

}

/*
How to tell zero values apart from assigned values
Zero values, although useful, can sometimes make it
diﬀicult to tell whether a given key has been assigned the
zero value, or if it has never been assigned.

Here’s an example of a program where this could be an
issue. This code erroneously reports that the student "Carl"
is failing, when in reality he just hasn’t had any grades
logged:
*/

func status1(name string) {
	// Declares a map using a map literal
	grades := map[string]float64{"peter": 0, "mary\n": 70}

	// A variable to hold a grade
	grade := grades[name]

	if grade < 60 {
		fmt.Printf("%s has failed\n", name)
	}

}

/*
To address situations, accessing a map key
optionally returns a second, Boolean value. It will be true if
the returned value has actually been assigned to the map,
or false if the returned value just represents the default
zero value.
*/
func status2(name string) {
	// Declares a map using a map literal
	grades := map[string]float64{"peter": 0, "mary\n": 70}

	// A variable to hold a grade
	grade, ok := grades[name]

	if !ok {
		fmt.Printf("No grade assigned to %s \n", name)

	} else if grade < 60 {
		fmt.Printf("%s has failed\n", name)
	}

}

func WordCount(s string) map[string]int {
	// Fields splits the string s around each instance of one or more consecutive white space characters,
	// as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains
	// only white space.
	sSlice := strings.Fields(s)

	counter := make(map[string]int)

	for _, item := range sSlice {
		counter[item]++
	}
	return counter
}
