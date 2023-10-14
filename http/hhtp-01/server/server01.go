package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a middleware function that reads the headers and cookies
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Retrieve and process the headers sent by the client
			userAgent := r.Header.Get("User-Agent")
			authorization := r.Header.Get("Authorization")

			// Use the headers as needed
			fmt.Printf("User Agent: %s\n", userAgent)
			fmt.Printf("Authorization: %s\n", authorization)

			// Get the cookie named "example_cookie"
			cookie, err := r.Cookie("example_cookie")
			if err != nil {
				if err == http.ErrNoCookie {
					// Cookie does not exist, so set a new one
					cookie = &http.Cookie{
						Name:  "example_cookie",
						Value: "Hello, Cookie!",
					}
					http.SetCookie(w, cookie)
				} else {
					// Some other error occurred
					fmt.Println(err.Error())
					return
				}
			}

			// Pass control to the next handler
			next(w, r)

			// Read the cookie value and send it in the response
			fmt.Fprintf(w, "Cookie Value: %s\n", cookie.Value)
		}
	}

	// Create the handler function that serves the main logic
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	}

	// Wrap the main handler with the middleware
	http.Handle("/", middleware(http.HandlerFunc(handler)))

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
