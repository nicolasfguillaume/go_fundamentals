// https://tour.golang.org/moretypes/22

package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Insert or update an element in map m
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// Retrieve an element
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Delete an element
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Test that a key is present with a two-value assignment
	// If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
