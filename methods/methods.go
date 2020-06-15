package methods

import (
	"image/color"
	"math"
	"sync"
)


var cache = struct {
	sync.Mutex
	mapping map[string]string
} {
	mapping: make(map[string]string),
}
func Lookup(key string) string {
	cache. Lock()
	v := cache.mapping[key]
	cache. Unlock()
	return v
}


type Point struct{ X, Y float64 }
// Traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
// The same, but as Point method
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


type Path []Point
// Distance returns path length
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}


type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Add(q Point) Point { return Point{p.X+q.X, p.Y+q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X-q.X, p.Y-q.Y} }

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call or path[i].Add(offset), or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
