package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Counter struct {
	a map[string]bool
	mu sync.Mutex
}

var i Counter = Counter{a: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher, waitG *sync.WaitGroup) {
	defer waitG.Done()
	if depth <= 0 {
		return
	}
	if i.checkvisited(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitG.Add(1)
		Crawl(u, depth-1, fetcher, waitG)
	}
	return
}

func (c Counter) checkvisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.a[url]
	if ok == false {
		c.a[url] = true
		return false
	}
	return true
}

func main() {
	var waitG sync.WaitGroup
	waitG.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &waitG)
	waitG.Wait()
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
