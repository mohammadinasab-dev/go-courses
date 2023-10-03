package main

import (
	"fmt"
)

type temporary interface {
	temp() bool
}

type MyDBConnectionError struct {
	Code    int
	Message string
}

func (e MyDBConnectionError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e MyDBConnectionError) temp() bool {
	return true
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.temp()
}

var maxRetry = 3

func main() {

	err := doSomething1()

	if err != nil {
		// Check if err is of type MyError
		if IsTemporary(err) {
			// Custom error handling for MyError type
			handleMyError(err)
		} else {
			// Generic error handling for other error types
			handleGenericError(err)
		}
	} else {
		fmt.Println("No errors!")
	}
}

func doSomething2() error {
	// Simulating an error occurrence.
	// return MyDBConnectionError{Code: 500, Message: "Something went wrong"}
	return nil
}

func doSomething1() error {
	// Simulating an error occurrence.
	return MyDBConnectionError{Code: 500, Message: "Something went wrong"}
	// return nil
}

func handleMyError(err error) {
	// Handle the error specific to MyError type
	fmt.Println("Handling Error based on behavior:", err)
	for i := 0; i < maxRetry; i++ {
		if err := doSomething2(); err != nil {
			fmt.Printf("retry %v\n", i)
			continue
		}
		fmt.Println("success")
		break
	}

	// Additional error handling logic for MyError can be added here.
}

func handleGenericError(err error) {
	// Handle generic error types
	fmt.Println("Handling generic error:", err)
	// Additional error handling logic for generic errors can be added here.
}
