package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// Note: No more need to store or check error return value when using panic
	// err := scanDirectory("./chapter2")

	//if err != nil {
	//log.Fatal(err)
	//}

	scanDirectory("chapter2")

}

/*
Each function that’s called - scanDirectory() - needs to return to the function
that called it. To enable this,Go keeps a call stack, a list of the function calls
that are active at any given point.

When a program panics, a stack trace, or listing of the call
stack, is included in the panic output. This can be useful in
determining what caused the program to crash.

Note:
When a program panics, all deferred function calls will still
be made. If there’s more than one deferred call, they’ll be
made in the reverse of the order they were deferred in.
*/

// When using panic - no more need to return error
func scanDirectory(path string) {
	// Print the current directory path
	fmt.Println(path)

	// ReadDir get a slice with the directory's content
	files, err := ioutil.ReadDir(path)

	if err != nil {
		// If an error value is returned from ReadDir ,
		// we pass it to panic instead.
		// Now, when scanDirectory encounters an error reading a
		// directory, it simply panics. All the recursive calls to
		// scanDirectory exit.

		panic(err)
	}

	// Join the directory path and file name with a slash
	for _, file := range files {
		filepath := filepath.Join(path, file.Name())

		// If this is a subdirectory
		if file.IsDir() {
			// Recursively call scanDirectory() this time with the subdirectory's path
			// Note: No more need to store or check error return value when using panic
			// err := scanDirectory(filepath)
			scanDirectory(filepath)

			/*
				if err != nil {
					// Print debug info from for error in ReadDir call
					fmt.Printf("Return error from scanDirectory(\"%s\") call\n", path)

					return err
				}
			*/

		} else {
			// Is this a regular file, just print it path
			fmt.Println(filepath)
		}

	}

}
