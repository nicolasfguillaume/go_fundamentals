// https://tour.golang.org/moretypes/24

package main

import (
	"fmt"
	"math"
)

// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.
// The compute function takes a function (that takes two float in argument,
// and returns a float) as argument and returns an float.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// inline definition of a function
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
