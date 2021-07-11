package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/stretchr/testify/assert"
)

func TestOptionsCorsResponse(t *testing.T) {
	// assert that OPTIONS request is properly served with correct CORS headers
	req := httptest.NewRequest(http.MethodOptions, "/url-analyze", strings.NewReader(`{"url": ""}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Accept, Content-Type, Content-Length, Authorization", w.Header().Get("Access-Control-Allow-Headers"))
}

func TestEmptyUrlField(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": ""}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	response, _ := ioutil.ReadAll(w.Result().Body)
	m := new(models.ErrorResponse)
	err := json.Unmarshal(response, m)

	// assert that correct responses are returned when url is empty
	assert.Nil(t, err)
	assert.Equal(t, 400, w.Result().StatusCode)
	assert.Equal(
		t, &models.ErrorResponse{ErrTitle: "validation error", ErrDescription: "url cannot be empty", ErrStatusCode: 400}, m)
}

func TestNoUrlField(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"test-url": "http://www.google.com"}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	response, _ := ioutil.ReadAll(w.Result().Body)
	m := new(models.ErrorResponse)
	err := json.Unmarshal(response, m)

	// assert that correct responses are returned when url is empty
	assert.Nil(t, err)
	assert.Equal(t, 400, w.Result().StatusCode)
	assert.Equal(
		t, &models.ErrorResponse{ErrTitle: "validation error", ErrDescription: "url cannot be empty", ErrStatusCode: 400}, m)
}

func TestInvalidUrlField(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": "invalid url"}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	response, _ := ioutil.ReadAll(w.Result().Body)
	m := new(models.ErrorResponse)
	err := json.Unmarshal(response, m)

	// assert that correct responses are returned when url is empty
	assert.Nil(t, err)
	assert.Equal(t, 400, w.Result().StatusCode)
	assert.Equal(
		t, &models.ErrorResponse{ErrTitle: "validation error", ErrDescription: "provided url is not valid", ErrStatusCode: 400}, m)
}

func TestUnreachableUrlField(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": "http://test@test.com"}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	response, _ := ioutil.ReadAll(w.Result().Body)
	m := new(models.ErrorResponse)
	err := json.Unmarshal(response, m)

	// assert that correct responses are returned when url is empty
	assert.Nil(t, err)
	assert.Equal(t, 400, w.Result().StatusCode)
	assert.Equal(
		t, &models.ErrorResponse{ErrTitle: "validation error", ErrDescription: "provided url is not reachable", ErrStatusCode: 400}, m)
}

func TestValidUrlAnalyzeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/url-analyze", strings.NewReader(`{"url": "http://example.com/"}`))
	w := httptest.NewRecorder()
	UrlAnalyzeHandler(w, req)
	response, _ := ioutil.ReadAll(w.Result().Body)
	m := new(models.HTMLPageDetails)
	err := json.Unmarshal(response, m)

	// assert that correct responses are returned when url is empty
	expected := &models.HTMLPageDetails{
		HTMLVersion: models.HTML5,
		Title:       "Example Domain",
		HeadingCount: models.HeadingCount{
			H1Count: 1,
			H2Count: 0,
			H3Count: 0,
			H4Count: 0,
			H5Count: 0,
			H6Count: 0,
		},
		LinkCount: models.LinkCount{
			InternalLinkCount:     0,
			ExternalLinkCount:     1,
			InaccessibleLinkCount: 0,
		},
		IsLoginForm: false,
	}
	assert.Nil(t, err)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, expected, m)
}
