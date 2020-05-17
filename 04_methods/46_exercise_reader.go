// https://tour.golang.org/methods/22

package main

// go get golang.org/x/tour/reader
import "golang.org/x/tour/reader"

// Implement a Reader type that emits an infinite stream 
// of the ASCII character 'A'

// The struct{} is a struct type with zero elements. It is often used 
// when no information is to be stored. It has the benefit of being 0-sized, 
// so usually no memory is required to store a value of type struct{}.
type MyReader struct{}

// Add a Read([]byte) (int, error) method to MyReader
func (r MyReader) Read(bytes []byte) (int, error) {
	var a byte = 'A' // ASCII code for A is 65 (one byte)
	for i := range bytes {
		bytes[i] = a
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
