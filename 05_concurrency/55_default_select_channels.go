// https://tour.golang.org/concurrency/6

package main

import (
	"fmt"
	"time"
)

func main() {
	// Provides access to the ticking channel
	// tick has type *<-chan time.Time
	// Sends to the tick channel every 100 ms
	tick := time.Tick(100 * time.Millisecond)
	// After reports whether the time instant t is after u
	// boom has type *<-chan time.Time
	// Sends to the boom channel when 500 ms has elasped from now
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		// Receives block when the buffer is empty
		case <-tick:
			fmt.Printf("tick.")
		// Receives block when the buffer is empty
		case <-boom:
			fmt.Println("BOOM!")
			return
		// The default case in a select is run if no other case is ready
		default:
			fmt.Println("    .")
			// Sleep pauses the program for at least 50 ms
			time.Sleep(50 * time.Millisecond)
		}
	}
}
