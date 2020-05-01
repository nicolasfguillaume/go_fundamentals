// https://tour.golang.org/basics/8

package main

import "fmt"

// Outside a function, every statement begins with a keyword (var, func, and so on)

// A var statement declares a list of variables; 
// as in function argument lists, the type is last
var c, python, java bool

// A var declaration can include initializers, one per variable
// If an initializer is present, the type can be omitted; 
// the variable will take the type of the initializer
var a, b = 1, 2

func main() {
	// Variables declared without an explicit initial value are given their zero value:
	// 	0 for numeric types, false for the boolean type, and "" for strings
	var i int
	var s string
	fmt.Println(i, s, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(a, b, c, python, java)

	// Inside a function, the := short assignment statement 
	// can be used in place of a var declaration with implicit type
	d, e := "NY", "NY"  // string must be in double quote (not single quote)
	fmt.Println(d, e)
}
