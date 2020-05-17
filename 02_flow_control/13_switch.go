// https://tour.golang.org/flowcontrol/9

 package main

import (
	"fmt"
	"time"
	"runtime"
)

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds
// Go only runs the selected case, not all the cases that follow

func main() {

	fmt.Print("Go runs on ")

	// Example 1
	switch os := runtime.GOOS; os {
		// if os == "darwin"
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")

	// Example 2
	today := time.Now().Weekday()
	switch time.Saturday {
		// if time.Saturday == today + 0
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tomorrow.")
		case today + 2:
			fmt.Println("In two days.")
		default:
			fmt.Println("Too far away.")
	}

	// Example 3
	// Switch without a condition is the same as switch true
	// This construct can be a clean way to write long if-then-else chains
	t := time.Now()
	switch {
		// if t.Hour() < 12
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
	}
}
