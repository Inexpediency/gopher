package surface

import (
	"fmt"
	"io"
	"math"
)

var sin30, cos30 = math.Sin(math.Pi / 6), math.Cos(math.Pi / 6) // sin(30°), cos(30°)

type Surf struct {
	Width, Height int     // canvas size in pixels
	Cells         int     // number of grid cells
	XYRange       float64 // axis ranges (-xyrange..+xyrange)
}

func (s *Surf) Draw(out io.Writer, f func(float64, float64) float64) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", s.Width, s.Height)

	xyscale := float64(s.Width) / 2 / s.XYRange // pixels per x or y unit
	zscale := float64(s.Height) * 0.4           // pixels per z unit

	for i := 0; i < s.Cells; i++ {
		for j := 0; j < s.Cells; j++ {
			ax, ay := corner(i+1, j, s, xyscale, zscale, f)
			bx, by := corner(i, j, s, xyscale, zscale, f)
			cx, cy := corner(i, j+1, s, xyscale, zscale, f)
			dx, dy := corner(i+1, j+1, s, xyscale, zscale, f)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, s *Surf, xyscale, zscale float64, f func(float64, float64) float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := s.XYRange * (float64(i)/float64(s.Cells) - 0.5)
	y := s.XYRange * (float64(j)/float64(s.Cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(s.Width)/2 + (x-y)*cos30*xyscale
	sy := float64(s.Height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func DefaultFunction(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
