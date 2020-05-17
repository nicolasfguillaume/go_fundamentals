// https://tour.golang.org/moretypes/23

/* 
Implement WordCount. It should return a map of the counts of each “word” 
in the string s. The wc.Test function runs a test suite against the 
provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

// go get golang.org/x/tour/wc
import (
	"strings"
	"golang.org/x/tour/wc"
)


func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range words {
		m[word] += 1
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
