package main

import (
	"fmt"
	"git_repo/go_playground/playground/variables/vars"
)

func main() {
	vars.BoilingPoint()

	fmt.Printf("%gF = %gC\n", vars.FreezingF, vars.FTOC(vars.FreezingF))
	fmt.Printf("%gC = %gF\n", vars.BoilingF, vars.FTOC(vars.BoilingF))

	//Test pointers
	vars.TestPointers()

}
