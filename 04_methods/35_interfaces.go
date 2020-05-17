// https://tour.golang.org/methods/9

package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures
type Abser interface {
	Abs() float64
}

/*
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit 
declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, 
which could then appear in any package without prearrangement. 
*/

type MyFloat float64

// Defining Abs method with value receiver (type MyFloat)
// This method means type MyFloat implements the interface Abser,
// but we don't need to explicitly declare that it does so.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Defining Abs method with pointer receiver (type *Vertex)
// This method means type *Vertex implements the interface Abser,
// but we don't need to explicitly declare that it does so.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// A value of interface type can hold any value that implements those methods
	// An interface value holds a value of a specific underlying concrete type
	var a Abser

	// math.Sqrt2 is constant = 1.4142135623730951
	// convert it to MyFloat type
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// Calling a method on an interface value executes the method of the same name 
	// on its underlying type:

	a = f  // a MyFloat type implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex type implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (the value type, not *Vertex)
	// and does NOT implement Abser because the Abs method is defined only on *Vertex 
	// (the pointer type).
	// a = v
}
