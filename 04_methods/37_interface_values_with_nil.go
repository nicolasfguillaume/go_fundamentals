// https://tour.golang.org/methods/12
// Interface values with nil underlying values

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	// A nil interface value holds neither value nor concrete type
	var i I

	// Calling a method on a nil interface is a run-time error because there is 
	// no type inside the interface tuple to indicate which concrete method to call.
	// describe(i)
	// i.M()

	var t *T
	// Note that an interface value that holds a nil concrete value is itself non-nil
	i = t            // a *T type implements I
	describe(i)
	// The concrete value inside the interface itself is nil, 
	// the method will be called with a nil receiver
	i.M()
	// In some languages this would trigger a null pointer exception, but in Go 
	// it is common to write methods that gracefully handle being called with 
	// a nil receiver

	i = &T{"hello"}  // a *T type implements I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
