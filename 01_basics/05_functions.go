// https://tour.golang.org/basics/4

package main

import "fmt"

// Notice that the type comes after the variable name
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters 
// share a type, you can omit the type from all but the last.
func sub(x, y int) int {
	return x - y
}

// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Return values may be named
// These names should be used to document the meaning of the return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// A return statement without arguments returns the named return values
	// This is known as a "naked" return
	return
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(sub(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}