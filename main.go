package main

import (
	"github.com/pandulaDW/home24-page-analyzer/handlers"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/url-analyze", handlers.UrlAnalyzeHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
