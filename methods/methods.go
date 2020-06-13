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
// Традиционная функция
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
// To же, но как метод типа Point
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


type Path []Point
// Distance возвращает длину пути,
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
