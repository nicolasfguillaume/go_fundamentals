// https://tour.golang.org/moretypes/13

package main

import "fmt"

// Slices can be created with the built-in make function; 
// this is how you create dynamically-sized arrays.
// The make function allocates a zeroed array and returns 
// a slice that refers to that array.
// Unlike new, make's return type is the same as the type 
// of its argument, not a pointer to it.

func main() {
	// returns a slice of len 5 (and cap 5)
	a := make([]int, 5)
	printSlice("a", a)

	// returns a slice of len 0 and cap 5
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("len=%d cap=%d %s=%v\n",
		len(x), cap(x), s, x)
}
