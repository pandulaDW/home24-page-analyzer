package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pandulaDW/home24-page-analyzer/handlers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/url-analyze", handlers.UrlAnalyzeHandler)

	// a random port will be assigned by heroku in production
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("server listening to requests at port %s...", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	d := time.Now().Add(30 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
