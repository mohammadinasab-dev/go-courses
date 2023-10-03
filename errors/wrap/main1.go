package main

// import (
// 	"errors"
// 	"fmt"
// )

// var (
// 	ErrNotFound    = errors.New("Not found error")
// 	ErrPermission  = errors.New("Permission denied")
// 	ErrInvalidData = errors.New("Invalid data")
// )

// func ReadData() error {
// 	// Simulating an error occurrence.
// 	originalErr := databaseQuery()

// 	// Wrapping the original error with additional context.
// 	err := fmt.Errorf("ReadData: %w", originalErr)
// 	return err
// }

// func databaseQuery() error {
// 	// Simulating an error occurrence.
// 	return ErrNotFound
// 	// return errors.New("#2054 relation violate over db under connections")
// }

// func main() {
// 	err := ReadData()

// 	fmt.Printf("err is: %v\n", err)
// 	if errors.Is(err, ErrNotFound) {
// 		fmt.Println("Handling Not Found error")
// 	} else if errors.Is(err, ErrPermission) {
// 		fmt.Println("Handling Permission error")
// 	} else if errors.Is(err, ErrInvalidData) {
// 		fmt.Println("Handling Invalid Data error")
// 	} else {
// 		fmt.Println("Handling generic error")
// 	}
// }
