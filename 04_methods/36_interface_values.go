// https://tour.golang.org/methods/11

package main

import (
	"fmt"
	"math"
)

/*
Under the hood, interface values can be thought of as a tuple of a value and a 
concrete type: (value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its
underlying type.
*/

type I interface {
	M()
}

// Defining method M with pointer type *T

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// Defining method M with value type F

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	// Calling a method on an interface value executes the method of the same name 
	// on its underlying type:

	i = &T{"Hello"}  // a *T type implements I
	describe(i)
	i.M()

	i = F(math.Pi)   // a F type implements I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
