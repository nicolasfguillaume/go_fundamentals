// https://tour.golang.org/methods/4

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// With a value receiver, the Scale method operates on a copy of the 
// original Vertex value. (This is the same behavior as for any other 
// function argument).
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Declaring methods with pointer receivers //
// Methods with pointer receivers can modify the value to which 
// the receiver points. Since methods often need to modify their 
// receiver, pointer receivers are more common than value receivers.
func (v *Vertex) Scale(f float64) {
	// Remember p.X is the same as (*p).X (as a convenience)
	v.X = v.X * f
	v.Y = v.Y * f
}
// The Scale method must have a pointer receiver 
// to change the Vertex value declared in the main function.

func main() {
	v := Vertex{3, 4}
	// update v in place
	v.Scale(10)
	fmt.Println(v.Abs())
}
