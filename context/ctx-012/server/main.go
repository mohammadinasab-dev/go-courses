package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Server starts handling the request...")
		ctx := req.Context()

		select {
		case <-ctx.Done():
			log.Println("Request has been canceled by the client")
			http.Error(w, "Request canceled", http.StatusGone)
			return
		case <-time.NewTimer(time.Second * 8).C:
			fmt.Fprintf(w, "Hi there!")
			log.Println("Request is done successfully")
			return
		case <-time.NewTimer(time.Second * 4).C:
			log.Println("Internal server error!")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

	})

	log.Println("Server listening on port 8085...")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal(err)
	}
}
