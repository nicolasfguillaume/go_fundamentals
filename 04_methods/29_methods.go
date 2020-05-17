// https://tour.golang.org/methods/1

package main

// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func 
// keyword and the method name.

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// The Abs method has a receiver of type Vertex named v
func (v Vertex) Abs_1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs written as a regular function with no change in functionality
func Abs_2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(v.Abs_1())
	fmt.Println(Abs_2(v))
}
