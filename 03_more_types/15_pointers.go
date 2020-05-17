// https://tour.golang.org/moretypes/1

package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value
// The type *T is a pointer to a T value. Its zero value is nil
// var p *int

func main() {
	i, j := 42, 2701

	// The & operator generates a pointer to its operand
	p := &i         // p is the memory address that points to i

	// The * operator denotes the pointer's underlying value (="dereferencing" or "indirecting")
	fmt.Println(*p) // read i through the pointer

	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // p is the memory address that points to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
