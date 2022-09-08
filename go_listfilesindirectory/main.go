package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

/*
The io/ioutil package includes a ReadDir function that will
let us read the directory contents. You pass ReadDir the
name of a directory, and it will return a slice of values, one
for each ﬁle or subdirectory the directory contains (along
with any error it encounters).
Each of the slice’s values satisﬁes the FileInfo interface,
which includes a Name method that returns the ﬁle’s name,
and an IsDir method that returns true if it’s a directory.
So our program calls ReadDir , passing it the name of
my_directory as an argument. It then loops over each value
in the slice it gets back. If IsDir returns true for the value, it
prints "Directory:" and the ﬁle’s name. Otherwise, it prints
"File:" and the ﬁle’s name.
*/

func main () {
	files, err := ioutil.ReadDir("my_directory")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			fmt.Println("Directory: ", file.Name())
			
		} else {
			fmt.Println("File: ", file.Name())

		}
		
	}
}