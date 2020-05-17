// https://tour.golang.org/methods/25

package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
	)

// Let's write another one, but this time it will return an implementation 
// of image.Image instead of a slice of data

const w = 600
const h = 400

func f(x, y int) uint8 {
	// Interesting functions include (x+y)/2, x*y, and x^y...
	return uint8((x*y)/2)
}

// Define your own Image type
type Image struct{}

// Implements the necessary methods for type Image:

// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h)
// Bounds returns the domain for which At can return non-zero color
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, w, h)
}

// ColorModel should return color.RGBAModel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the color of the pixel at (x, y)
func (i Image) At(x, y int) color.Color {
	// the value v in the last picture generator corresponds 
	// to color.RGBA{v, v, 255, 255} in this one
	v := f(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
