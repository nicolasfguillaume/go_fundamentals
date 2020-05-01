// https://tour.golang.org/flowcontrol/5

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// concat strings
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Like for, the if statement can start with a 
	// short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		// Variables declared by the statement are only in 
		// scope until the end of the if.
		return v
	} else {
		// Variables declared inside an if short statement 
		// are also available inside any of the else blocks.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		sqrt(2), 
		sqrt(-4),
	)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
