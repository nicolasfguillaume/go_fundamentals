// https://tour.golang.org/methods/14

package main

import "fmt"

func main() {
	// The interface type that specifies zero methods is known as the empty interface
	var i interface{}
	describe(i)

	// An empty interface may hold values of any type (Every type implements at 
	// least zero methods)
	// Empty interfaces are used by code that handles values of unknown type

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
