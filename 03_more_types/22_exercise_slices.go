// https://tour.golang.org/moretypes/18

/* Implement Pic. It should return a slice of length dy, each element of which 
is a slice of dx 8-bit unsigned integers. When you run the program, it will 
display your picture, interpreting the integers as grayscale 
(well, bluescale) values.*/


package main

// to install the package:
// go get golang.org/x/tour/pic
import "golang.org/x/tour/pic"

func f(x, y int) uint8 {
	// Interesting functions include (x+y)/2, x*y, and x^y...
	return uint8((x+y)/2)
}

func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)

	for y := range picture {
		picture[y] = make([]uint8, dx)

		for x := range picture[y] {
			picture[y][x] = f(x, y)
		}

	}

	return picture
}

func main() {
	// creates the image and encodes it as base64
	pic.Show(Pic)
}
