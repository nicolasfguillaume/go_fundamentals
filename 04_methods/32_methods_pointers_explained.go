// https://tour.golang.org/methods/5

package main

import (
	"fmt"
	// "math"
)

type Vertex struct {
	X, Y float64
}

// Scale method
// Methods with value receivers take either a value or a pointer 
// as the receiver when they are called
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale method rewritten as a function
// Functions that take a value argument must take a value of that specific type
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// &v has type *Vertex
	ScaleFunc(&v, 10)  // &v is the memory address that points to v
	fmt.Println(v)

	// Nethods with pointer receivers take either a value or a pointer 
	// as the receiver when they are called
	/* For the statement v.Scale(5), even though v is a value and not a 
	pointer, the method with the pointer receiver is called automatically. 
	That is, as a convenience, Go interprets the statement v.Scale(5) 
	as (&v).Scale(5) since the Scale method has a pointer receiver.*/
	v = Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)

	p := &Vertex{3, 4}  // p is the memory address that points to Vertex{3, 4}
	p.Scale(10)
	fmt.Println(*p)     // read Vertex{3, 4} through the pointer
}
