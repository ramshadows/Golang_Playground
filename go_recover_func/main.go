package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
Go oﬀers a built-in recover function that can stop a program
from panicking. Use it to exit the program gracefully.

Note: When you call recover during normal program execution, it
just returns nil and does nothing else:

      *************************************
	  func main() {
		  // When you call "recover" in a program that isn't panicking
		  // it doea nothing and returns nil

		  fmt.PrintIn(recover())  //prints nil
	  }

Note2: If you call recover when a program is panicking, it will stop
the panic. But when you call panic in a function, that
function stops executing. So there’s no point calling recover
in the same function as panic , because the panic will
continue anyway:

       *********************************************
	   func freakOut(){
		   // The panic stops the rest of the code below it from running
		   panic("Oh no")

		   recover()  // Never runs
	   }

	   fun main() {
		   freakOut() // Program crashes here - stack trace printed

		   fmt.PrintIn("Exiting normally") // never runs
	   }

Note3: But there is a way to call recover when a program is
panicking... During a panic, any deferred function calls are
completed. So you can place a call to recover in a separate
function, and use defer to call that function before the code
that panics.

*********************************************

       func calmDown(){
		   // Call "recover()" in this other function
		   recover()
	   }
	   func freakOut(){
		   // Defer a call to a function that recovers
		   defer calmDown()

		   // The program panics after that - deferred function will recover
		   // Any code after the panic statement never runs
		   panic("Oh no")

		   fmt.PrintIn("Never runs")
	   }

	   fun main() {
		   freakOut() // Program does not crashes here - so stack trace is never printed

		   fmt.PrintIn("Exiting normally") // runs
	   }

Note4: But when there is a panic, recover returns whatever value
was passed to panic . This can be used to gather information
about the panic, to aid in recovering or to report errors to
the user.

*********************************************

       func calmDown(){
		   // Call "recover" and print the panic value
		   fmt.PrintIn(recover())
	   }

	   fun main() {
		   defer calmDown()
		   panic("Oh no") // This is the value that will be returned from "recover"
	   }
*/

func main() {
	// You only call recover in a function where you're also calling defer. 
	// If you call the recover() function, it returns nil and has no other effect in normal running.
	// Defer a call to our reportPanic() before calling
	// code that might panic

	defer reportPanic()

	//panic("Some unanticipated panic")

	scanDirectory("./chapter2")

}

func reportPanic() {

	// Call recover and store its return value
	p := recover()

	// If recover() returned nil - there's no panic
	if p == nil {
		// so do nothing
		return
	}

	// Otherwise get the underlying error value
	// Note: We achive this by doing a Type Assertion
	// We assert that the type of the panic value is an error
	// ie: we take the return value of recover and converts it back
	// to an error value. Once that’s done, we can safely call the Error method.

	err, ok := p.(error)

	if ok {
		// If that conversion is successful, we print the
		// error value.
		fmt.Println(err)

	} else {
		// Re-instate the panic to handle unanticipated panics
		// since the anticipated panic was not an error
		panic(p)
	}

}

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
