package surface

import (
	"io"
	"math"

	"github.com/ajstarks/svgo"
)

// Surf struct
type Surf struct {
	Width   int
	Height  int
	Cells   int
	XYrange float64
	XYscale float64
	Zscale  float64
	Angle   float64
}

// Draw surface
func Draw(out io.Writer, d *Surf) {

	s := svg.New(out)
	s.Start(d.Width, d.Height)
	s.Style("text/css", "stroke: red; fill: white; stroke-Width: 0.7;")

	for i := 0; i < d.Cells; i++ {
		for j := 0; j < d.Cells; j++ {
			ax, ay := corner(i+1, j, d)
			bx, by := corner(i, j, d)
			cx, cy := corner(i, j+1, d)
			dx, dy := corner(i+1, j+1, d)

			x := make([]int, 0)
			y := make([]int, 0)
			x = append(x, int(ax), int(bx), int(cx), int(dx))
			y = append(y, int(ay), int(by), int(cy), int(dy))

			s.Polygon(x, y, "stroke: grey; fill: white; stroke-Width: 0.7;")
		}
	}

	s.End()
}

func corner(i, j int, d *Surf) (float64, float64) {
	var sin, cos = math.Sin(d.Angle), math.Cos(d.Angle)

	x := d.XYrange * (float64(i) / float64(d.Cells) - 0.5)
	y := d.XYrange * (float64(j) / float64(d.Cells) - 0.5)

	z := f(x, y)

	sx := float64(d.Width/2) + (x*y)*cos*d.XYscale
	sy := float64(d.Height/2) + (x+y)*sin*d.XYscale - z*d.Zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}
