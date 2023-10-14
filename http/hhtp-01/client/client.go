package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:8080/pattern1"

	// Create a new request with headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Add headers to the request
	req.Header.Add("User-Agent", "MyCustomUserAgent")
	req.Header.Add("Authorization", "Bearer your-token-here")

	// Send the modified request
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		cookies := response.Cookies()
		for _, cookie := range cookies {
			fmt.Printf("Cookie Name: %s\n", cookie.Name)
			fmt.Printf("Cookie Value: %s\n", cookie.Value)
		}

		body := make([]byte, 1024)
		_, err = response.Body.Read(body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(body))
	} else {
		fmt.Println("Error: Unexpected status code received")
	}
}
