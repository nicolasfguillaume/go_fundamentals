// https://tour.golang.org/methods/21

package main

import (
	"fmt"
	"io"
	"strings"
)

/* 
The io package specifies the io.Reader interface, which represents 
the read end of a stream of data. The Go standard library contains many 
implementations of these interfaces, including files, network connections, 
compressors, ciphers, and others.

The io.Reader interface has a Read method:
func (T) Read(b []byte) (n int, err error) 
*/


func main() {
	// Create a strings.Reader (that implements the Read method)
	r := strings.NewReader("Hello, Reader!")

	// Then consume its output 8 bytes at a time:

	// Make a slice of len 8 (= transfer buffer)
	b := make([]byte, 8)
	for {
		// Read populates the given byte slice with data and returns 
		// the number of bytes populated and an error value. 
		// It returns an io.EOF error when the stream ends.
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		// Convert bytes to string (to double-quote strings, use %q)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
