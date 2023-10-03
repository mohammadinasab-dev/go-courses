package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	// TLS configuration
	cert, err := tls.LoadX509KeyPair("../certificate.crt", "../private.key")
	if err != nil {
		log.Fatal(err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Create TCP listener
	listener, err := tls.Listen("tcp", ":8085", &config)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on :8085")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Handle incoming connections
	request := make([]byte, 1024)
	n, err := conn.Read(request)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Received message: %s", string(request[:n]))

	response := []byte("Hello from the server!")
	_, err = conn.Write(response)
	if err != nil {
		log.Println(err)
		return
	}
}
