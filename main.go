package main

import (
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"context"
)

func main() {
	// Create custom logger
	logger := log.New(os.Stdout, "products-api", log.LstdFlags)
	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)

	// Create a dedicated Serve Mux
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Create a dedicated server so we can fine tune our requirements
	server := &http.Server{
		Addr:              ":9090",
		Handler:           sm,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
	}

	// Encapsulate into a go func to avoid blocking
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Initialize a buffered channel of type os.Signal
	sigChan := make(chan os.Signal)
	// Listen for interrupt and kill os signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Forward signal into channel
	sig := <-sigChan
	logger.Println("Received Terminate, gracefully shutting down", sig)

	// Graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
