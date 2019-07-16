package svg_surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 900, 480            // grid size
	cells         = 100                 // Count of cells on the grid
	xyrange       = 30.0                // Axis range
	xyscale       = width / 2 / xyrange // count of pixels in one x or y
	zscale        = height * 0.1        // count of pixels in one z
	angle         = math.Pi / 6         // Angle of axis x, y (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Draw(out io.Writer) {
	generateHeader(out)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			color := determineColor(ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Fprintf(out, "<polygon fill='%s' points='%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func determineColor(ax, ay, bx, by, cx, cy, dx, dy float64) string {
	c := "blue"
	// bottom part of canvas
	if by-dy > 0.001 {
		c = "black"
	}
	return c
}

func generateHeader(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
}

func corner(i, j int) (float64, float64) {
	// calc angle point (x, y) of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// calc hight of surface z
	z := f(x, y)
	// project (x, y, z) on SVG canvas (sx, sy) in isometric mode
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from 0, 0
	return math.Sin(r)
}
