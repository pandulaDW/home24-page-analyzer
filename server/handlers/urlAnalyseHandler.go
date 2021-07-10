package handlers

import (
	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/pandulaDW/home24-page-analyzer/services"
	"net/http"
)

// UrlAnalyzeHandler will handle the requests coming to the url-analyze route
func UrlAnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	// set content type header
	w.Header().Set("Content-Type", "application/json")

	// decode and validate request body
	requestBody, err := models.DecodeRequestURLBody(r)
	if err != nil {
		models.EncodeError(w, "validation error", err.Error(), http.StatusBadRequest)
		return
	}

	// validate and get the url content
	requestURL, err := http.Get(requestBody.Url)
	if err != nil || requestURL.StatusCode != http.StatusOK {
		models.EncodeError(w, "validation error", "provided url is not valid", http.StatusBadRequest)
		return
	}

	// get the service response
	response := services.HtmlPageDetails(requestURL.Body)

	// encode the service response
	models.EncodeResponseUrlAnalyse(w, response)
}
