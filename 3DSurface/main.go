package main

import (
	"fmt"
	"git_repo/go_playground/playground/3DSurface/surface"
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", surface.Width, surface.Height)
	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {
			ax, ay := surface.Corner(i+1, j)
			bx, by := surface.Corner(i, j)
			cx, cy := surface.Corner(i, j+1)
			dx, dy := surface.Corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
