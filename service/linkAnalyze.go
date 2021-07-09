package service

import (
	"fmt"
	"net/http"
	"sync"
)

// checkAccessibility checks if each of the external links in the linkCountObj
// are accessible and returns the sum of inaccessible links.
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
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(link)
		return false
	}
	return true
}
