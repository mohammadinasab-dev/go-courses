package main

import (
	"errors"
	"fmt"
	"log"
)

/*
	- Errors are values.

	- Errors are not exceptional situations
	  but rather ordinary values that can be handled gracefully within the code.

	- Sentinel errors" refer to specific values that
	  are used to indicate certain conditions or outcomes in a program.

	- Sentinel errors are not special from a compiler perspective.
	  Rather, they are a way of conceptualizing certain, specific,
	  errors that the developer wants to handle in specific ways because they signal specific things.

	- "Comparing the string form of an error is,
	   in my opinion, a code smell, and you should try to avoid it. by Dave Cheney"

	- The common practice for error handling relies on guard clauses(early return).

	- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
*/

var (
	ErrNotFound = errors.New("not found")
	ErrTooLarge = errors.New("too large")
)

func main() {

	if err := getFromDB(); err == ErrNotFound {
		log.Println(err)
		fmt.Println("insert into db")
	}
}

func getFromDB() error {
	return ErrNotFound
}

func WriteInto() error {
	return ErrTooLarge
}
