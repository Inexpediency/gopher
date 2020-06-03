package malbedro

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
)

func Draw(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	err := png.Encode(out, img)
	if err != nil {
		log.Fatal("Error encoding image")
	}
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast = 20
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255-contrast*n}
		}
	}

	return color.Black
}