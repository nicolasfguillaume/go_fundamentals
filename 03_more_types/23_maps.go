// https://tour.golang.org/moretypes/19

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// a map maps keys to values (like a dictionary in Python)
// The zero value of a map is nil. 
// A nil map has no keys, nor can keys be added.
var m map[string]Vertex

func main() {
	// The make function returns a map of the given type 
	// (key: string ; value: Vertex), 
	// initialized and ready for use.
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals are like struct literals, but the keys are required.
	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	// If the top-level type is just a type name, you can omit it 
	// from the elements of the literal.
	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)

}
