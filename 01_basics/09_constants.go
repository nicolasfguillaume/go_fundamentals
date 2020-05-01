// https://tour.golang.org/basics/15

package main

import "fmt"

// Constants cannot be declared using the := syntax
const Pi = 3.14

// Create a huge number by shifting a 1 bit left 100 places.
// In other words, the binary number that is 1 followed by 100 zeroes.
const Big = 1 << 100

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
