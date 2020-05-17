// https://tour.golang.org/methods/3

package main

import (
	"fmt"
	"math"
)

// Declaring a non-struct type that derives from float64 type
type MyFloat float64

// Declaring a method on non-struct types //
// You can only declare a method with a receiver whose type is 
// defined in the same package as the method. You cannot declare 
// a method with a receiver whose type is defined in another 
// package (which includes the built-in types such as int).
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// math.Sqrt2 is constant = 1.4142135623730951
	// convert it to MyFloat type
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
