// https://tour.golang.org/basics/13

package main

import (
	"fmt"
	"math"
)

/* Assignment between items of different type requires an explicit conversion
Some numeric conversions:
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
or put more simply:
i := 42
f := float64(i)
u := uint(f)
*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	// When declaring a variable without specifying an explicit type,
	// the variable's type is inferred from the value on the right hand side.
	v := 42.0 
	fmt.Printf("v is of type %T\n", v)
}
