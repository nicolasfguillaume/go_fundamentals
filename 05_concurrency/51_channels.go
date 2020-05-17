// https://tour.golang.org/concurrency/2

package main

import "fmt"

/*
The example code sums the numbers in a slice, distributing the 
work between two goroutines. Once both goroutines have completed 
their computation, it calculates the final result.
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Sends sum to channel c
	c <- sum 
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Channels are a typed conduit through which you can send and 
	// receive values with the channel operator, <-.
	// Like maps and slices, channels must be created before use
	c := make(chan int)

	// Starts new goroutines running sum(s, c)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// Receives from channel c and assign value to x and y
	// (The data flows in the direction of the arrow)
	x, y := <-c, <-c 
	// By default, sends and receives block until the other side is ready. 
	// This allows goroutines to synchronize without explicit locks or 
	// condition variables.

	fmt.Println(x, y, x+y)
}
