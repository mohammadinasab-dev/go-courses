package main

import (
	"errors"
	"fmt"
)

/*
	- It works with wrapped errors,
	  where an error wraps another error using the fmt.Errorf, errors.Wrap, or similar functions.
*/

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {
	err := doSomething()

	if err != nil {
		// Check if err matches MyError
		if errors.Is(err, MyError{Code: 500, Message: "Something went wrong"}) {
			// Custom error handling for MyError type
			handleMyError(err)
		} else {
			// Generic error handling
			handleGenericError(err)
		}
	} else {
		fmt.Println("No errors!")
	}
}

func doSomething() error {
	// Simulating an error occurrence.
	return MyError{Code: 500, Message: "Something went wrong"}
}

func handleMyError(err error) {
	// Handle the error specific to MyError type
	fmt.Println("Handling MyError:", err)
	// Additional error handling logic for MyError can be added here.
}

func handleGenericError(err error) {
	// Handle generic error types
	fmt.Println("Handling generic error:", err)
	// Additional error handling logic for generic errors can be added here.
}
