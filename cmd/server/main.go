/*
Copyright © 2023 Caezarr-OSS

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Caezarr-OSS/test-api/api"
)

var version = "0.1.0"

func main() {
	log.Printf("Starting test-api server version %s", version)
	
	// Create the API router
	router := api.NewRouter()
	
	// Define server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	
	// Start server in a goroutine
	go func() {
		log.Printf("Server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	
	// Setup graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	// Wait for termination signal
	<-stop
	
	log.Println("Shutting down server...")
	fmt.Println("Server stopped")
}
