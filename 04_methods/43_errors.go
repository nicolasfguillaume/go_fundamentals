// https://tour.golang.org/methods/19

package main

import (
	"fmt"
	"time"
	"strconv"
)

/* Go programs express error state with error values.
The error type is a built-in interface similar to fmt.Stringer:
    type error interface {
        Error() string
    }
(As with fmt.Stringer, the fmt package looks for the error interface 
when printing values.) */

type MyError struct {
	When time.Time
	What string
}

// Implementation of a custom error message
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// Make the function run() returns a type error
// The fmt package looks for the error interface when printing value of the error
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	// Functions often return an error value, and calling code should handle 
	// errors by testing whether the error equals nil.
	i, err := strconv.Atoi("42")
	if err != nil {
	    fmt.Printf("couldn't convert number: %v\n", err)
	    return
	}
	fmt.Println("Converted integer:", i)
	// A nil error denotes success; a non-nil error denotes failure.
	fmt.Println("Error of conversion:", err)
	
	// Calling the implementation of a custom error message
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
