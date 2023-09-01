package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

type result struct {
	px, py int
	z      complex128
}

func main() {
	wg := &sync.WaitGroup{}
	result_chan := make(chan result)
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 10000, 10000
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		wg.Add(1)
		go func(py int, result_chan chan result) {
			for px := 0; px < width; px++ {

				y := float64(py)/height*(ymax-ymin) + ymin
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				result_chan <- result{px, py, z}

			}
			wg.Done()
		}(py, result_chan)
		// Image point (px, py) represents complex value z.

	}
	go func(img *image.RGBA) {
		for result := range result_chan {
			img.Set(result.px, result.py, sqrt(result.z))
		}
	}(img)
	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
