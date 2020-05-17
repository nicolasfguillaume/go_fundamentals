// https://tour.golang.org/concurrency/9

package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutual exclusion (conventional name for the data structure that provides it is
mutex): if we don't need communication amoung goroutines, if we just want to make
sure only one goroutine can access a variable at a time to avoid conflicts.

Go provides mutual exclusion with sync.Mutex and its two methods: 
Lock and Unlock
*/

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	// Go's dictionnary type
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// Block of code to be executed in mutual exclusion by surrounding it with 
	// a call to Lock and Unlock
	// The method will stop and must wait that other goroutines had unlocked 
	// the variable, before being able to access/modify it.
	// Here accessing the method Lock() of the mux field of variable c, in order to
	// lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()

	// We can also use defer to ensure the mutex will be unlocked.
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
