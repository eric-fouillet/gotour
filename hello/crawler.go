package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

/*
 * A special struct that holds extra crawling info
 */
type ParallelCrawler struct {
	LocalFetcher Fetcher
	Visited      map[string]bool
	MapMutex     chan bool
	PrinterMutex chan bool
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *ParallelCrawler) Crawl(url string, depth int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	c.MapMutex <- true
	c.Visited[url] = true
	<-c.MapMutex
	body, urls, err := c.LocalFetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.PrinterMutex <- true
	fmt.Printf("found: %s %q\n", url, body)
	<-c.PrinterMutex
	for _, u := range urls {
		c.MapMutex <- true
		isVisited := c.Visited[url]
		<-c.MapMutex
		if !isVisited {
			done := make(chan bool, 1)
			done <- true
			go c.Crawl(u, depth-1)
			<-done
		}
	}
	return
}

/*func CrawlParallel(url string, depth int, fetch Fetcher) {
	m := map[string]bool{url: true}
	var crawler func(url string, depth int, fetch Fetcher)
	crawler = func(url string, depth int, fetch Fetcher) {
		if depth <= 0 {
			return
		}
		if m[url] {
			return
		}
		m[url] = true
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
	}
}*/

func main() {
	c := &ParallelCrawler{fetcher, make(map[string]bool), make(chan bool, 1), make(chan bool, 1)}
	c.Crawl("http://golang.org/", 4)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
