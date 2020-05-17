// https://tour.golang.org/concurrency/5

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// The select statement lets a goroutine wait on multiple 
		// communication operations: a select blocks until one of 
		// its cases can run, then it executes that case. 
		// It chooses one at random if multiple are ready.
		select {
		// Sends to a channel
		case c <- x:
			x, y = y, x+y
		// Receives block when the buffer is empty
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// Initializes the channels
	c := make(chan int)
	quit := make(chan int)

	// Starts an new (anonymous) goroutine to populate the channels
	// By default, sends and receives block until the other side is ready. 
	// This allows goroutines to synchronize without explicit locks or 
	// condition variables.
	go func() {
		for i := 0; i < 10; i++ {
			// Receives block when the buffer is empty
			fmt.Println(<-c)
		}
		// Sends to a channel
		quit <- 0
	}()

	// Passes the channels to the function
	fibonacci(c, quit)
}
