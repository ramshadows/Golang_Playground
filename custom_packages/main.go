package main

import (
	"fmt"
	"git_repo/go_playground/playground/custom_packages/tempcomputation"
	"git_repo/go_playground/playground/variables/vars"
)

func main() {
	fmt.Println(tempcomputation.CToF(tempcomputation.BoilingC))

	fmt.Println(tempcomputation.CToF(vars.FreezingF))

}
