// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var ch0 chan ResultObject
// var lockStatus = 0

// type ResultObject struct {
// 	Result fakeResult
// 	Err    error
// }

// func contains[T comparable](s []T, e T) bool {
// 	for _, a := range s {
// 		if a == e {
// 			return true
// 		}
// 	}
// 	return false
// }

// type Fetcher interface {
// 	// Fetch returns the body of URL and
// 	// a slice of URLs found on that page.
// 	Fetch(url string, ch chan ResultObject)

// }

// // Crawl uses fetcher to recursively crawl
// // pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher, ch chan ResultObject) {

// 	// TODO: Fetch URLs in parallel.
// 	// TODO: Don't fetch the same URL twice.
// 	// This implementation doesn't do either:
// 	fetcher.mu.Lock()
// 	lockStatus = 1
// 	if depth <= 0 {
// 		return
// 	}
// 	_, fetchedBefore := fetcher.fetchedUrls[url]
// 	if !fetchedBefore {
// 		fetcher.fetchedUrls[url] = true
// 		go fetcher.Fetch(url, ch)
// 		resultObj := <-ch
// 		if resultObj.Err != nil {
// 			fmt.Println(resultObj.Err)
// 			return
// 		}
// 		fmt.Printf("found: %s %q\n", url, resultObj.Result.body)
// 		for _, u := range resultObj.Result.urls {
// 			if lockStatus == 1 {
// 				fetcher.mu.Unlock()
// 				lockStatus = 0
// 			}
// 			Crawl(u, depth-1, fetcher, ch)
// 		}
// 	}

// 	defer unlock(&fetcher.mu)
// 	return
// }

// func unlock(mu *sync.Mutex) {
// 	if lockStatus == 1 {
// 		mu.Unlock()
// 		lockStatus = 0
// 	}
// }
// func main() {
// 	ch0 = make(chan ResultObject)
// 	fetcher.fetchedUrls = make(map[string]bool)
// 	Crawl("https://golang.org/", 4, &fetcher, ch0)
// }

// // FakeFatcher is Fetcher that returns canned results.
// type FakeFatcher struct {
// 	mu          sync.Mutex
// 	fakeResults map[string]*fakeResult
// 	fetchedUrls map[string]bool
// }

// type fakeResult struct {
// 	body string
// 	urls []string
// }

// func (f *FakeFatcher) Fetch(url string, ch chan ResultObject) {
// 	if res, ok := f.fakeResults[url]; ok {
// 		ch <- ResultObject{fakeResult{res.body, res.urls}, nil}
// 	} else {
// 		ch <- ResultObject{fakeResult{}, fmt.Errorf("not found: %s", url)}
// 	}
// }

// // fetcher is a populated FakeFatcher.
// var fetcher = FakeFatcher{
// 	fakeResults: map[string]*fakeResult{
// 		"https://golang.org/": &fakeResult{
// 			"The Go Programming Language",
// 			[]string{
// 				"https://golang.org/pkg/",
// 				"https://golang.org/cmd/",
// 			},
// 		},
// 		"https://golang.org/pkg/": &fakeResult{
// 			"Packages",
// 			[]string{
// 				"https://golang.org/",
// 				"https://golang.org/cmd/",
// 				"https://golang.org/pkg/fmt/",
// 				"https://golang.org/pkg/os/",
// 			},
// 		},
// 		"https://golang.org/pkg/fmt/": &fakeResult{
// 			"Package fmt",
// 			[]string{
// 				"https://golang.org/",
// 				"https://golang.org/pkg/",
// 			},
// 		},
// 		"https://golang.org/pkg/os/": &fakeResult{
// 			"Package os",
// 			[]string{
// 				"https://golang.org/",
// 				"https://golang.org/pkg/",
// 			},
// 		},
// 	},
// }
