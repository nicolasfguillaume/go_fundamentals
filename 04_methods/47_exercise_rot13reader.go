// https://tour.golang.org/methods/23

package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

/*
A common pattern is an io.Reader that wraps another io.Reader, 
modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader 
(a stream of compressed data) and returns a *gzip.Reader that 
also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads 
from an io.Reader, modifying the stream by applying the rot13 
substitution cipher to all alphabetical characters.
*/

// The rot13Reader type is provided. Make it an io.Reader 
// by implementing its Read method.
type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
		// rot13 on majuscules
		case b >= 'A' && b <= 'M':
			b += 13
		case b > 'M' && b <= 'Z':
			b -= 13
		// rot13 on minuscules
		case b >= 'a' && b <= 'm':
			b += 13
		case b > 'm' && b <= 'z':
			b -= 13
	}

	return b
}

// Adding a Read method to rot13Reader
func (r13 rot13Reader) Read(bytes []byte) (int, error) {
	// Because there are 26 letters (2×13) in the basic Latin alphabet, 
	// ROT13 is its own inverse; that is, to undo ROT13, the same algorithm 
	// is applied, so the same action can be used for encoding and decoding.

	// Read populates the given byte slice with data and returns 
	// the number of bytes populated and an error value
	// Remmber that struct fields are accessed using a dot: r13.r
	n, err := r13.r.Read(bytes)
	// Modifying each byte in place
	for i := range bytes {
		bytes[i] = rot13(bytes[i])
	}

	return n, err
}

func main() {
	// Create a strings.Reader (that implements the io.Read method)
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// Returns a struct containing an io.Reader
	r := rot13Reader{s}
	// Copies the content of in-memory reader r and copies it to stdout
	io.Copy(os.Stdout, &r)
}
