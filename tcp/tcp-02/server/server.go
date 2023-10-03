package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8085"
	TYPE = "tcp"
)

/*
	What would happen if we didn't close the listener?

	If you don't close the listener in your code, it will continue to listen for incoming connections indefinitely,
	even after the main function exits. This means that the TCP socket associated with the listener will remain open, occupying system resources.

	Leaving the listener open can lead to problems such as resource leakage and potential denial of service if too many connections are allowed to accumulate.
	Additionally, if you try to rerun the program,
	while the previous instance is still running and holding onto the same port, you may encounter an error indicating that the port is already in use.
*/

/*
	What would happen if we didn't close the connections?

	Resource consumption: Each open connection consumes system resources, including memory and network buffers.
	If you have a large number of unclosed connections, it can lead to resource exhaustion and degrade the performance of your application.

	Connection limits: Operating systems often have limits on the maximum number of open connections that can be simultaneously maintained.
	Failing to close connections can eventually exceed these limits, causing new connection attempts to be rejected.

	Scalability issues: If your application doesn't release connections promptly,
	it may not be able to handle new connection requests efficiently. This can impact its ability to scale and serve a large number of clients concurrently.

	Connection leaks: Unclosed connections can lead to connection leaks, where resources associated with the connection are not properly released.
	This can cause memory leaks and result in instability or crashes in long-running applications.
*/

func main() {
	log.Println("Starting TCP server!")

	listen, err := net.Listen(TYPE, fmt.Sprintf(HOST+":"+PORT))
	handleServerErrors(err)

	// Catch error and close listener when main function exits
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleRequests(conn)
	}

}

func handleRequests(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	handleReadWriteErrors(err)
	log.Printf("Message fro client is: %v\n", string(buffer))

	response := fmt.Sprintf("here i am the server")

	_, err = conn.Write([]byte(response))
	handleReadWriteErrors(err)

}

func handleServerErrors(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func handleReadWriteErrors(err error) {
	if err != nil {
		log.Println("Error writing / reading response:", err)
	}
}
