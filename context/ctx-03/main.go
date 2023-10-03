package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start a goroutine to listen for the cancellation signal.
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		fmt.Println("Received shutdown signal...")
		cancel()
	}()

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
			cancel()
		}
		// go handleFunction(ctx)
	}()

	// Perform other operations
	performOperation(ctx)

	// Wait for the context to be canceled
	<-ctx.Done()

	// Perform cleanup or any final tasks before exiting
	fmt.Println("Shutting down gracefully...")
	time.Sleep(2 * time.Second) // Simulating cleanup tasks

	// Close the server
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("Error shutting down server:", err)
	}

	fmt.Println("Application shutdown complete.")
}

func performOperation(ctx context.Context) {
	// Perform the operation here
	// ...
	for {
		select {
		case <-ctx.Done():
			// Handle cancellation if needed
			fmt.Println("Operation canceled")
			return
		default:
			// Continue performing the operation
			fmt.Println("Performing operation...")
			time.Sleep(time.Second)
		}
	}
}