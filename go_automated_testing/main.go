package main

import (
	"fmt"
	joinstrings "git_repo/go_playground/playground/go_automated_testing/joinStrings"
)

func main() {
	phrases := []string{"Mangoes", "Oranges", "Pears", "Bananas"}

	fmt.Println(joinstrings.JoinWithCommas(phrases))
	fmt.Println("A photo of ", joinstrings.JoinWithCommas(phrases))

}
