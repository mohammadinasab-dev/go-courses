package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Load the CA certificate
	caCert, err := ioutil.ReadFile("../certificate.crt")
	if err != nil {
		log.Fatal(err)
	}

	// Create the certificate pool
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// TLS configuration
	config := &tls.Config{
		InsecureSkipVerify: true,
		// RootCAs: caCertPool,
	}

	// Dial TCP connection with TLS configuration
	conn, err := tls.Dial("tcp", "127.0.0.1:8085", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	message := []byte("Hello from the client!")
	_, err = conn.Write(message)
	if err != nil {
		log.Println(err)
		return
	}

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Received message: %s\n", string(response[:n]))
}
