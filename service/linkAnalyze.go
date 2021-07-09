package service

import (
	"net/http"
	"sync"
	"time"
)

// checkAccessibility checks if each of the external links in the linkCountObj
// are accessible and returns the sum of inaccessible links.
//
// If the links take more than 3 seconds to get a response code, it will be considered
// as an inaccessible link.
func checkAccessibility(externalLinks []string) int {
	ch := make(chan bool)
	var inAccessibleCount int

	go func() {
		var wg sync.WaitGroup
		defer close(ch)

		for _, link := range externalLinks {
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				for {
					select {
					case ch <- isAccessible(link):
						return
					case <-time.After(3 * time.Second):
						ch <- false
						return
					}
				}
			}(link)
		}
		wg.Wait()
	}()

	for val := range ch {
		if !val {
			inAccessibleCount++
		}
	}

	return inAccessibleCount
}

// isAccessible returns true if the link is accessible, false otherwise.
//
// If the links returns a http.StatusOK response for a GET request, it
// will be considered as an accessible link. If it throws a protocol error
// or returns a status code which is not http.StatusOK, it will be considered as an inaccessible link.
func isAccessible(link string) bool {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
