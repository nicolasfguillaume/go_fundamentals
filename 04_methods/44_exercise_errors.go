// https://tour.golang.org/methods/20

package main

import (
	"fmt"
	"math"
)

// Create a new type
type ErrNegativeSqrt float64

// And make it an error by giving it a method
// ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
func (e ErrNegativeSqrt) Error() string {
	// A call to fmt.Sprint(e) inside the Error method will send the program 
	// into an infinite loop. Avoid this by converting e first to float
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Copy your Sqrt function from the earlier exercise and modify it 
// to return an error value. Sqrt should return a non-nil error value 
// when given a negative number, as it doesn't support complex numbers.
// Change your Sqrt function to return an ErrNegativeSqrt value 
// when given a negative number.
func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	eps := 1E-6

	// anonymous function
	f := func(x, z float64) float64 { 
    	return - (z*z - x) / (2*z)
  	} 

	for math.Abs(f(x, z)) > eps {
		z += f(x, z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
