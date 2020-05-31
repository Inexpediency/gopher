package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x41, 0x69, 0xE1, 0xFF}, color.RGBA{0xE1, 0x73, 0xB1, 0xff}, color.RGBA{0xCE, 0xFF, 0x00, 0xFF}}

// Draw lissajous GIF image
func Draw(io io.Writer, cycles int, res float64, size, nframes, delay int) {
	lissajous(io, cycles, res, size, nframes, delay)
}

func peekColor(palette []color.Color) int {
	return rand.Int()%(len(palette)-1) + 1
}

func lissajous(out io.Writer, cycles int, res float64, size, nframes, delay int) {
	// const (
	// 	cycles  = 5     // Number of vibrations
	// 	res     = 0.001 // Angular resolution
	// 	size    = 500   // The image canvas covers [size..+size]
	// 	nframes = 128   // The number of animation frames
	// 	delay   = 5     // Delay between frames
	// )

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Relative frequency of vibrations
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase defference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(peekColor(palette)))
		}

		phase += 0.05
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
