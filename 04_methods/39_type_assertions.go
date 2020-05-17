// https://tour.golang.org/methods/15

package main

import "fmt"

func main() {
	var i interface{} 
	i = "hello"

	// A type assertion provides access to an interface value's 
	// underlying concrete value

	// This statement asserts that the interface value i holds the concrete type T 
	// and assigns the underlying T value to the variable t
	// t := i.(T)

	s := i.(string)
	fmt.Println(s)

	// To test whether an interface value holds a specific type, a type assertion 
	// can return two values: the underlying value and a boolean value that reports 
	// whether the assertion succeeded
	// If i holds a T, then t will be the underlying value and ok will be true
	s, ok := i.(string)
	fmt.Println(s, ok)

	// If not, ok will be false and t will be the zero value of type T, 
	// and no panic occurs
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// If i does not hold a float64, this statement will trigger a panic.
	// f = i.(float64) // panic
	// fmt.Println(f)
}
