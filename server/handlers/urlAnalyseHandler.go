package handlers

import (
	"log"
	"net/http"

	"github.com/pandulaDW/home24-page-analyzer/models"
	"github.com/pandulaDW/home24-page-analyzer/services"
)

// setupCorsResponse write the cors orgin headers to the response
func setupCorsResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

// UrlAnalyzeHandler will handle the requests coming to the url-analyze route
func UrlAnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(http.MethodGet + " " + r.URL.Path)

	// set content type and cors headers
	setupCorsResponse(w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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
