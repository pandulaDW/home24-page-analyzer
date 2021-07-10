package models

import (
	"encoding/json"
	"errors"
	"net/http"
)

// RequestUrlAnalyse defines the structure of the request body for url analyse route
type RequestUrlAnalyse struct {
	Url string `json:"url"`
}

// DecodeRequestURLBody decodes the data from the given request body's input stream and identify
// whether it's conforming to RequestUrlAnalyse structure.
//
// Returns error if it doesn't correspond or the if there's a decoding issue
func DecodeRequestURLBody(r *http.Request) (*RequestUrlAnalyse, error) {
	decoder := json.NewDecoder(r.Body)
	requestBody := new(RequestUrlAnalyse)
	err := decoder.Decode(requestBody)
	if err != nil {
		return nil, err
	}

	if requestBody.Url == "" {
		return nil, errors.New("url cannot be empty")
	}

	return requestBody, nil
}

// ResponseUrlAnalyse defines the structure of the response body for url analyse route
type ResponseUrlAnalyse struct {
	HTMLPageDetails
}

// EncodeResponseUrlAnalyse encodes the given HTMLPageDetails data to the response writer's output stream.
// The function will send a 500 a internal server error, if encoding is unsuccessful.
func EncodeResponseUrlAnalyse(w http.ResponseWriter, pageDetails *HTMLPageDetails) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusOK)

	encodeError := encoder.Encode(pageDetails)
	if encodeError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

// ErrorResponse defines the standard error response sent to the client
type ErrorResponse struct {
	ErrTitle       string `json:"err_message"`
	ErrDescription string `json:"err_description"`
	ErrStatusCode  int    `json:"err_status_code"`
}

// EncodeError encodes the given error data to the response writer's output stream.
// The function will send a 500 a internal server error, if encoding is unsuccessful.
func EncodeError(w http.ResponseWriter, err string, errDescription string, statusCode int) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(statusCode)

	errBody := ErrorResponse{
		ErrTitle:       err,
		ErrDescription: errDescription,
		ErrStatusCode:  statusCode,
	}

	encodingErr := encoder.Encode(errBody)
	if encodingErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
