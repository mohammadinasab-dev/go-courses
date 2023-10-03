package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
remove for of Accept.

remove go from handle

remove for in handle

remove closing the listener and connection
*/

func main() {

	fmt.Println("server is starting over 8085...")
	ln, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		con, err := ln.Accept()
		log.Printf("connection accepted %#v\n", con)
		if err != nil {
			log.Fatal(err)
		}

		go handle(con)
	}

	// time.Sleep(time.Second * 3)
	// log.Println("server terminated")

}

// func handle(con net.Conn) {

// 	// defer func() {
// 	// 	if err := con.Close(); err != nil {
// 	// 		log.Println("error closing connection: ", err)
// 	// 	}
// 	// }()

// 	defer con.Close()

// 	scanner := bufio.NewScanner(con)

// 	for scanner.Scan() {
// 		data := scanner.Text()
// 		fmt.Println(data)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Println(err)
// 	}
// }

func handle(con net.Conn) {
	defer con.Close()

	// for i := 0; i < 3; i++ {
		data, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		// sth to do here ...
		fmt.Println(data)
		// return 
	// }

}

// handle Just One Line
// func handle(con net.Conn) {
// 	// defer con.Close()

// 	scanner := bufio.NewScanner(con)

// 	if scanner.Scan() {
// 		data := scanner.Text()
// 		fmt.Println(data)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Println(err)
// 	}
// 	// time.Sleep(time.Second * 9)

// }
