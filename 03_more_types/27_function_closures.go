// https://tour.golang.org/moretypes/25

package main

import "fmt"

// Go functions may be closures. A closure is a function value that references 
// variables from outside its body. The function may access and assign to the 
// referenced variables; in this sense the function is "bound" to the variables.

// The adder function returns a closure (a function which takes int as argument 
// and returns an int). Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// The adder function returns a closure function
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
