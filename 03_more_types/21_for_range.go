// https://tour.golang.org/moretypes/16

package main

import "fmt"

// The range form of the for loop iterates over a slice or map.

func main() {
	var pow1 = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// When ranging over a slice, two values are returned for each 
	// iteration. The first is the index, and the second is a copy 
	// of the element at that index.
	for i, v := range pow1 {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)

	// If you only want the index, you can omit the second variable v.
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}

	// You can skip the index or value by assigning to _.
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
