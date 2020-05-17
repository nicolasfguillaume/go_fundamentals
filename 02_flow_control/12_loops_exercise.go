// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

// Objective: implement a square root function
func Sqrt(x float64) float64 {
	z := 1.0
	eps := 1E-6

	// anonymous function
	f := func(x, z float64) float64 { 
    	return - (z*z - x) / (2*z)
  	} 

	for math.Abs(f(x, z)) > eps {
		z += f(x, z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
