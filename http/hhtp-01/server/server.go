package main

import (
	"fmt"
	"net/http"
)

// type MyHandler struct{}

// func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello"))
// }

// func main() {
// 	handler := MyHandler{}

// 	http.Handle("/", handler)

// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	http.Handle("/hello", http.HandlerFunc(helloHandler))
// 	http.HandleFunc("/world", worldHandler)

// 	http.ListenAndServe(":8080", nil)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello"))
// }

// func worldHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("World"))
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

		// Read the cookie value and send it in the response
		fmt.Fprintf(w, "Cookie Value: %s", cookie.Value)
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
