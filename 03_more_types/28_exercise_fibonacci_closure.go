// https://tour.golang.org/moretypes/26

package main

import "fmt"

// Implement a fibonacci function that returns a function (a closure) 
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	
	return func() int {

		res := a 
		// fibonacci
		c := a + b
		// translation
		a = b
		b = c

		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
