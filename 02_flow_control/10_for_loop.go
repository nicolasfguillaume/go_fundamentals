// https://tour.golang.org/flowcontrol/1

package main

import "fmt"

// Go has only one looping construct, the for loop.
// the init statement (optional): executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement (optional): executed at the end of every iteration

func main() {
	
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// The init and post statements are optional (you can drop the ;)
	// C's while is spelled for in Go
	sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	// without the loop condition, it loops forever
	// for {
	// }
}
