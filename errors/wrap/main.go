package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

var (
	ErrNotFound    = errors.New("Not found error")
	ErrPermission  = errors.New("Permission denied")
	ErrInvalidData = errors.New("Invalid data")
)

func ReadData() error {
	// Simulating an error occurrence.
	originalErr := databaseQuery()

	// Wrapping the original error with additional context.
	err := fmt.Errorf("ReadData: %w", originalErr)
	return err
}

func main() {
	err := ReadData()

	if errors.Is(err, MyError{
		Code: 401,
		Message: "notFound",
		}) {
		fmt.Println("Handling Not Found error")
	} else if errors.Is(err, ErrPermission) {
		fmt.Println("Handling Permission error")
	} else if errors.Is(err, ErrInvalidData) {
		fmt.Println("Handling Invalid Data error")
	} else {
		fmt.Println("Handling generic error")
	}
}

func databaseQuery() error {
	// Simulating an error occurrence.
	return MyError{
		Code:    401,
		Message: "notFound",
	}
}
