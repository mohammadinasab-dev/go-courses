package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//TODO: Test and Complete it

func main() {
	// Create a listener to accept incoming connections
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// Create a context with cancellation mechanism
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a wait group to track active connections
	var wg sync.WaitGroup

	// Create a channel to receive OS signals for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine to handle shutdown signals
	go func() {
		<-signalCh
		fmt.Println("\nReceived shutdown signal...")
		cancel()
	}()

	// Start accepting and handling connections
	fmt.Println("TCP server started, listening on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			// Check if the error is a result of closing the listener
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				fmt.Println("Temporary accept error:", err)
				time.Sleep(5 * time.Second)
				continue
			}
			break // Break the loop on permanent errors
		}

		// Handle the connection in a separate goroutine
		wg.Add(1)
		go handleConnection(ctx, conn, &wg)
	}

	// Stop accepting new connections
	listener.Close()

	// Wait for all active connections to finish
	wg.Wait()

	fmt.Println("TCP server shut down gracefully.")
}

func handleConnection(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	// Notify the wait group when the connection handling is done
	defer wg.Done()

	// Start a goroutine to monitor the connection status
	go func() {
		select {
		case <-ctx.Done():
			// Perform any necessary cleanup before closing the connection
			fmt.Println("Connection closed due to server shutdown")
		}
	}()

	// Handle the connection here
	// ...

	// Wait for the shutdown signal or connection closure
	select {
	case <-ctx.Done():
		// Perform any necessary cleanup before closing the connection
		fmt.Println("Connection closed due to server shutdown")
	}

	// Close the connection
	conn.Close()
}
