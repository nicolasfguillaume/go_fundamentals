// https://tour.golang.org/methods/7

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// The equivalent thing happens in the reverse direction.

// Functions that take a value argument must take a value of that specific type
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// While methods with value receivers take either a value or a pointer 
// as the receiver when they are called
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &v
	// The method call p.Abs() is interpreted as (*p).Abs(). It can be either a value or pointer.
	fmt.Println(p.Abs())
	// Here, the argument must be a value (not a pointer)
	fmt.Println(AbsFunc(*p))
}
