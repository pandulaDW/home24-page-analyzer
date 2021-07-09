package service

import (
	"net/http"
	"sync"
	"time"
)

var (
	// httpClient defines a http client with the specified timeout
	httpClient = http.Client{Timeout: 5 * time.Second}
)

// checkAccessibility checks if each of the external links in the linkCountObj
// are accessible and returns the sum of inaccessible links.
//
// If a link takes more than 5 seconds to get a response, it will be considered as inaccessible
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
				ch <- isAccessible(link)
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
	resp, err := httpClient.Get(link)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
