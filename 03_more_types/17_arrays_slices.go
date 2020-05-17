// https://tour.golang.org/moretypes/6

package main

import "fmt"

// Arrays //
/* The type [n]T is an array of n values of type T:
   var a [2]string
   An array's length is part of its type, so arrays cannot be resized */

// Slices //
/* The type []T is a slice with elements of type T:
   var b []string
   A slice is dynamically-sized
   a[low : high]
   This selects a half-open range which includes the first element,
   but excludes the last one.
   Note: slices are like references to arrays (a slice does not store any data,
   it just describes a section of an underlying array) --> Changing the elements
   of a slice modifies the corresponding elements of its underlying array.
   The length and capacity of a slice s can be obtained using len(s) and cap(s)
   The length of a slice is the number of elements it contains.
   The capacity of a slice is the number of elements in the underlying array, 
   counting from the first element in the slice */

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// defining an array literal
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// new slice of this array
	var s []int = primes[1:4]
	fmt.Printf("len=%d cap=%d s=%v\n", len(s), cap(s), s)

	// The default is zero for the low bound
	s = s[:2]
	fmt.Printf("len=%d cap=%d s=%v\n", len(s), cap(s), s)

	// The default is the length of the slice for the high bound
	s = s[1:]
	fmt.Printf("len=%d cap=%d s=%v\n", len(s), cap(s), s)

	// slice litteral (like an array literal without the length)
	// creates an array [3]bool{true, true, false}
	// then builds a slice that references it:
	r := []bool{true, true, false}
	fmt.Println(r)

	// The zero value of a slice is nil (a nil slice has a length 
	// and capacity of 0 and has no underlying array):
	var n []int
	if n == nil {
		fmt.Println("nil!")
	}

	// slice of struct
	d := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(d)

	// changing the elements of a slice modifies the corresponding
	// elements of its underlying array
	s[0] = 1000
	fmt.Println(primes)
}
