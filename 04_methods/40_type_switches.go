// https://tour.golang.org/methods/16

package main

import "fmt"

func do(i interface{}) {
	// A type switch is a construct that permits several type assertions in series
	// The cases in a type switch specify types (not values), and those values are 
	// compared against the type of the value held by the given interface value

	// The declaration in a type switch has the same syntax as a type assertion i.(T), 
	// but the specific type T is replaced with the keyword type.

	// v hold the value held by i, and the type of the underlying concrete value
	switch v := i.(type) {
		// if v has type int
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		// if v has type int
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		// no match; here v has the same type as i
		default:
			fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
