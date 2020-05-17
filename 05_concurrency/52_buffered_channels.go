// https://tour.golang.org/concurrency/3

package main

import "fmt"

func main() {
	// Channels can be buffered. Provide the buffer length as the second 
	// argument to make to initialize a buffered channel
	ch := make(chan int, 2)

	// Sends to a buffered channel block only when the buffer is full
	ch <- 1
	ch <- 2

	// Receives block when the buffer is empty
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
