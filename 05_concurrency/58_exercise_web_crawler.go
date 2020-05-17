// https://tour.golang.org/concurrency/10
// Use Go's concurrency features to parallelize a web crawler.

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *SafeCounter) Crawl(url string, depth int, fetcher Fetcher) {

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.

	if depth <= 0 {
		return
	}

	// if url to be fetched was not already fetched, then fetch it
	// if url not in c.v {
	// 	...
	// }

	body, urls, err := go fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		// add url to cache
		//c.mux.Lock()
		// c.v[url] = ...
		c.Crawl(u, depth-1, fetcher)
		//c.mux.Unlock()
	}
	return

}

/*
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}
*/

func main() {
	c := SafeCache{v: make(map[string][]string])}
	c.Crawl("https://golang.org/", 4, fetcher)
}

// Hint: you can keep a cache of the URLs that have been fetched on a map, 
// but maps alone are not safe for concurrent use!
// SafeCounter is safe to use concurrently.
type SafeCache struct {
	// mapping an url to urls
	v   map[string][]string
	mux sync.Mutex
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
