package main

import (
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8085"
	TYPE = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Println("Unable to resolve the server: ", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Println("Unable to make connection to the server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("here i am the client"))
	if err != nil {
		log.Println("Unable to write into the connection: ", err)
		os.Exit(1)
	}

	receiver := make([]byte, 1024)
	// conn.Close()
	_, err = conn.Read(receiver)
	if err != nil {
		log.Println("Unable to read from the connection: ", err)
		os.Exit(1)
	}
	log.Printf("Received message from the server is: %v", string(receiver))

}
