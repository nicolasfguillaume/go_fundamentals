// https://tour.golang.org/basics/1

package main

// By convention, the package name is the same as the last element of the import path
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

// Note: The environment in which these programs are executed is deterministic, 
// so each time you run the example program rand.Intn will return the same number
