package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	err := scanDirectory("./chapter2")

	if err != nil {
		log.Fatal(err)
	}

}

/*
The scanDirectory function takes the path of the directory it should scan,
so we pass it the path to scan.
The ﬁrst thing scanDirectory does is print the current path,
so we know what directory we’re working in. Then it calls
ioutil.ReadDir on that path, to get the directory contents.

It loops over the slice of FileInfo values that ReadDir
returns, processing each one. It calls filepath.Join to join
the current directory path and the current ﬁlename
together with slashes

If the current ﬁle isn’t a directory, scanDirectory just prints
its full path, and moves on to the next ﬁle (if there are any
more in the current directory).
But if the current ﬁle is a directory, the recursion kicks in:
scanDirectory calls itself with the subdirectory’s path. If that
subdirectory has any subdirectories, scanDirectory will call
itself with each of those subdirectories, and so on through
the whole ﬁle tree.
*/

func scanDirectory(path string) error {
	// Print the current directory path
	fmt.Println(path)

	// ReadDir get a slice with the directory's content
	files, err := ioutil.ReadDir(path)

	if err != nil {
		// Print debug info from for error in ReadDir call
		fmt.Printf("Return error from scanDirectory(\"%s\") call\n", path)

		return err
	}

	// Join the directory path and file name with a slash
	for _, file := range files {
		filepath := filepath.Join(path, file.Name())

		// If this is a subdirectory
		if file.IsDir() {
			// Recursively call scanDirectory() this time with the subdirectory's path
			err := scanDirectory(filepath)

			if err != nil {
				// Print debug info from for error in ReadDir call
				fmt.Printf("Return error from scanDirectory(\"%s\") call\n", path)
				
				return err
			}

		} else {
			// Is this a regular file, just print it path
			fmt.Println(filepath)
		}

	}

	return nil

}
