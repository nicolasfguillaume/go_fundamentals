// https://tour.golang.org/methods/24

package main

import (
	"fmt"
	"image"
)

/*
Package image defines the Image interface (https://golang.org/pkg/image/#Image):

type Image interface {
    ColorModel() color.Model   // color.Model type is also an interface (predefined implementation is color.RGBAModel)
    Bounds() Rectangle         // Rectangle type is actually an image.Rectangle type
    At(x, y int) color.Color   // color.Color type is also an interface (predefined implementation is color.RGBA)
}

Package color defines the following interfaces:

type Color interface {
    RGBA() (r, g, b, a uint32)
}

type Model interface {
    Convert(c Color) Color
}

*/

func main() {
	// image package has the function Rect(x0, y0, x1, y1 int) Rectangle
	// A Rectangle is also an Image whose bounds are the rectangle itself
	// NewNRGBA returns a new NRGBA image (type *image.NRGBA) with the given bounds
	// NRGBA is an in-memory image whose At method returns color.NRGBA values
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Applies method func (p *RGBA) Bounds() Rectangle
	fmt.Println(m.Bounds())

	// Applies method func (p *RGBA) At(x, y int) color.NRGBA
	// then applies method func (c NRGBA) RGBA() (r, g, b, a uint32)
	fmt.Println(m.At(0, 0).RGBA())
}
