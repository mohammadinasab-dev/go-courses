package main

// import (
// 	"errors"
// 	"fmt"
// )

// /*
// 	- The ability to wrap an underlying error to provide more context!
// */

// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

// func (e *PathError) Error() string {
// 	return fmt.Sprintf("%s %s: %v", e.Op, e.Path, e.Err)
// }

// func OpenFile(path string) error {
// 	// Simulating an error occurrence.
// 	innerErr := errors.New("failed to open file")

// 	// Wrapping the error in a PathError.
// 	err := &PathError{
// 		Op:   "open",
// 		Path: path,
// 		Err:  innerErr,
// 	}

// 	return err
// }

// func main() {
// 	filePath := "/path/to/file.txt"
// 	err := OpenFile(filePath)

// 	if pathErr, ok := err.(*PathError); ok {
// 		fmt.Printf("Error occurred while %s file '%s': %v\n", pathErr.Op, pathErr.Path, pathErr.Err)
// 	} else {
// 		fmt.Println("Unknown error occurred.")
// 	}
// }
