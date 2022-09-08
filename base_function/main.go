/*
#basename() Function

Removes any prefix of string that looks like a file system path with
components separated by slashes, and it removes any suffix that looks like a file type

#Example
fmt.Println(basename("a/b/c.go")) // "c"
fmt.Println(basename("c.d.go"))  // "c.d"
fmt.Println(basename("abc"))  // "abc"
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Using Function baseName1: ", baseName1("a/b/c.go"))

	fmt.Println("Using Function baseName2: ", baseName2("a/b/c.go"))

}

// The first version of basename does all the work without the help of libraries:
func baseName1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("i = %d value of s[i] = %q\n", i, s[i])
		if s[i] == '/' {
			s = s[i+1:]
			fmt.Printf("Value of s = %s \n", s)
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("i = %d value of s[i] = %q\n", i, s[i])
		if s[i] == '.' {
			s = s[:i]
			fmt.Printf("Value of s = %s \n", s)
			break
		}
	}
	return s
}

// A simpler version uses the strings.LastIndex librar y func tion:
func baseName2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
