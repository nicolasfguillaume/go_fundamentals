// https://tour.golang.org/concurrency/1

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime
	// go f(x, y, z) starts a new goroutine running f(x, y, z)
	// The evaluation of f, x, y, and z happens in the current goroutine 
	// and the execution of f happens in the new goroutine.
	go say("world")
	// Goroutines run in the same address space, so access to shared memory 
	// must be synchronized. The sync package provides useful primitives, 
	// although you won't need them much in Go as there are other primitives.

	say("hello")
}
