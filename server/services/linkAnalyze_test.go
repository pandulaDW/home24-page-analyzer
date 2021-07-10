package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestIsAccessible(t *testing.T) {
	// assert that broken links are correctly returned
	assert.False(t, isAccessible("https://mrdoob.github.com/three.js"))

	// assert that malformed links are correctly returned
	assert.False(t, isAccessible("some test link"))

	// assert that correct links are correctly returned
	assert.True(t, isAccessible("https://www.google.com"))

	// assert that timeout works correctly
	CurrentHttpClient := httpClient
	httpClient = http.Client{Timeout: 50 * time.Millisecond}
	assert.False(t, isAccessible("https://www.amazon.com"))
	httpClient = CurrentHttpClient
}

func TestCheckAccessibility(t *testing.T) {
	// assert that correct inaccessible count is returned
	links := []string{"https://www.google.com", "http://example.com", "https://mrdoob.github.com/three.js"}
	assert.Equal(t, 1, checkAccessibility(links))
}
