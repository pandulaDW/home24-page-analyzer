package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlAnalyzeHandler(t *testing.T) {
	// assert that OPTIONS request is properly served
	req := httptest.NewRequest(http.MethodOptions, "/url-analyze", strings.NewReader(`{"url": ""}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	// assert that correct responses are returned when url is empty
	req = httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": ""}`))
	w = httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	// assert that correct responses are returned when url field is not present
	req = httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"test-url": "http://www.google.com"}`))
	w = httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	// assert that correct responses are returned when url is not valid
	req = httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": "http://www.google"}`))
	w = httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	// assert that correct responses are returned when url is not valid
	req = httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": "http://www.google"}`))
	w = httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 400, w.Result().StatusCode)

	// assert that correct responses are returned when url is valid
	req = httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": "https://www.google.com"}`))
	w = httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}
