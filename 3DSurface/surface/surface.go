// Surface computes an SVG rendering of a 3-D surface function.

package surface

import "math"

const (
	Width, Height = 600, 320
	Cells         = 100
	Xyrange       = 30.0
	Xyscale       = Width / 2 / Xyrange
	Zscale        = Height * 0.04
	Angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(Angle), math.Cos(Angle) // sin(30°), cos(30°)

func Corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := Xyrange * (float64(i)/Cells - 0.5)
	y := Xyrange * (float64(j)/Cells - 0.5)
	// Compute surface height z.
	z := F(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*cos30*Xyscale
	sy := Height/2 + (x+y)*sin30*Xyscale - z*Zscale
	return sx, sy
}
func F(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
