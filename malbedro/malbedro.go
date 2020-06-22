package malbedro

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
)

func DrawMalbedro(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
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
		contrast   = 20
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}

	return color.Black
}

func DrawNuton(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		eps                    = 0.1
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			if x != 0 || y != 0 {

				t := z
				z = 0.8*z + 0.2*cmplx.Pow(z, -4)

				for cmplx.Abs(z-t) >= eps {
					t = z
					z = 0.8*z + 0.2*cmplx.Pow(z, -4)
				}

				var cl color.Color
				cl = color.Black
				r := int(cmplx.Phase(z) / (math.Pi / 5))
				if r == 0 {
					cl = color.RGBA{R: 0xE1, G: 0x39, B: 0x58, A: 0xFF}
				} else if r == 2 || r == 1 {
					cl = color.RGBA{R: 0xB0, G: 0xE1, B: 0x80, A: 0xFF}
				} else if r == 4 || r == 3 {
					cl = color.RGBA{R: 0xFF, G: 0xE1, B: 0x00, A: 0xFF}
				} else if r == -4 || r == -3 {
					cl = color.RGBA{R: 0x58, G: 0x6F, B: 0xE1, A: 0xFF}
				} else if r == -2 || r == -1 {
					cl = color.RGBA{R: 0xD7, G: 0x1F, B: 0xE1, A: 0xFF}
				}

				img.Set(px, py, cl)
			}
		}
	}

	err := png.Encode(out, img)
	if err != nil {
		log.Fatal("Error encoding image")
	}
}
