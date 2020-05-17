// https://tour.golang.org/methods/17

package main

import "fmt"
/* One of the most ubiquitous interfaces is Stringer defined by the fmt package
  type Stringer interface {
      String() string
  }
A Stringer is a type that can describe itself as a string. The fmt package 
(and many others) look for this interface to print values */

type Person struct {
	Name string
	Age  int
}

// Add a "String() string" method to the type Person
// This makes the Person type implement the interface fmt.Stringer 
// to print with format "name (age years)"
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
