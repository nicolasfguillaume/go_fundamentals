// https://tour.golang.org/moretypes/2

package main

import "fmt"

// A struct is a collection of fields
// A struct type is a schema containing the blueprint of a data that
// a structure will hold. We need to create a new derived type 
// so that we can refer to the struct type. We use struct keyword 
// to create a new structure type.
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1 = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	// Struct fields are accessed using a dot
	v.X = 4
	fmt.Println(v.X)

	v = Vertex{1, 2}
	// Struct fields can be accessed through a struct pointer
	p := &v  // p is the memory address that points to v
	p.X = 1e9  // p.X is the same as (*p).X (as a convenience)
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)
}
