package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
	- In certain server designs, a separate thread is created for each incoming network request.
	  This enables multiple requests to be processed simultaneously using multiple threads.
	  Each thread handles either a single request or a group of related requests.

	- Newer approach is that servers using the non-blocking I/O, event loops, or callbacks techniques.

	 "Non-blocking I/O is a way for a program to keep doing other things without waiting for I/O operations to finish.
	  It doesn't get stuck waiting for reading or writing data, so it can work on other tasks in the meantime.
	  This makes the program more efficient and responsive because it can handle multiple tasks at the same time."

	 "The event loop is in charge of checking different sources for events, like user input, network activity, timers, and files.
	  When it finds one that's ready, it calls the appropriate callback or handler for that event.
	  By using callbacks and non-blocking I/O, the event loop allows things to happen simultaneously without one task blocking another.
	  It's all about being asynchronous and super efficient!"

	  "nginx works using an event-driven architecture technique, based on an event loop."

	- Goroutines in Go programming language enable us to utilize event looping and non-blocking I/O effectively.

	- Internally, the Accept() function uses blocking system calls provided by the operating system to wait for incoming connections.

	- A file descriptor is the numerical identifier used by an operating system to uniquely identify an open network connection.
	  in Unix-like systems, network programming functions such as socket(), accept(), read(), and write() work with file descriptors.

	- We almost close the connection, not the listeners.

	- TLS stands for Transport Layer Security. It is a cryptographic protocol that provides secure communication over a computer network.
	  It operates at the transport layer of the network protocol stack and
	  establishes an encrypted connection between two parties, allowing them to exchange data securely.

	- The stream-like nature of TCP means TCP provides a reliable, ordered, and byte-oriented data transmission service.
	  It treats data as a continuous stream of bytes rather than individual packets or messages.

	- Almost everything in the Golang net package is synchronous, and it's the magic of net package.
*/

func main() {

	ln, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go CopyToStdOutWithTimeOut(conn)
	}
}

func CopyToStdOut(conn net.Conn) {
	n, err := io.Copy(os.Stdout, conn)
	log.Printf("Copied %d bytes; finished wit err = %v", n, err)
}

func CopyToStdOutWithTimeOut(conn net.Conn) {
	defer conn.Close() // explain it
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("finished wit err = %v", err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
}

func CopyToStdOutWithTimeOutWithErr(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("Read timeout")
				continue
			} else {
				log.Printf("finished with err = %v", err)
				return
			}
		}
		os.Stderr.Write(buf[:n])
	}
}

func handleConnectionBySize(conn net.Conn) {
	defer conn.Close()

	// Size in bytes to read
	size := 1024

	buf := make([]byte, size)
	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println("Error reading from connection:", err)
		}
		return
	}

	fmt.Printf("Received %d bytes: %s\n", n, buf[:n])
}
