// https://tour.golang.org/methods/8

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/* There are two reasons to use a pointer receiver (over a value receiver):
- first: the method can modify the value that its receiver points to
- second: to avoid copying the value on each method call. This can be more efficient 
if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex,
even though the Abs method needn't modify its receiver.

All methods on a given type should have either value or pointer receivers, 
but not a mixture of both.

*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}  // p is the memory address that points to Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
