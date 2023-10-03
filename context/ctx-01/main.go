package main

import (
	"context"
	"fmt"
	"time"
)

/*

   - A context can be used to propagate cancellation signals across goroutines or
     to pass values and deadlines down the call chain without explicitly passing them as function arguments.

   - It is a common practice to pass a context as the first argument in function signatures
     when working with concurrent or long-running operations.
     e.g. networking, HTTP requests, database interactions, and other I/O-bound operations

   - It is indeed a common practice to create a child context instead of directly using the parent context.

   - Avoid storing large objects, such as database connections or complete data structures, in the context.
     Instead, use the context to pass smaller values like request IDs, API keys, user authentication details,
     or other necessary metadata.

   - The context package provides a way to carry request-scoped values across API boundaries and goroutines.
	 do not pass it as a struct fields.

	- To safely retrieve a value stored in the context without risking any panic or failure due to the interface type of the value,
	  you can write a function that performs type assertion.

    - It is necessary to defer the cancel function of the context even if we use context.WithTimeout or WithDeadLine.

	- Never start a goroutine unless you are aware of the condition under which it should be closed.


*/

func main() {
	// Create a context with cancellation capability
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	// Start a goroutine to perform some operation
	go performOperationGraceFully(ctx)
	// go performOperation(ctx)

	// Another operation which is related to the "performOperationGraceFully"
	func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()

	// Another operation
	func() {
		// time.Sleep(time.Second * 3)
		<-ctx.Done()
	}()

	fmt.Println("Main goroutine terminated.")
}

func performOperationGraceFully(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		// Check if the context has been canceled
		case <-ctx.Done():
			//sth to do ...
			fmt.Println("Operation gracefully closed.")
			return
		default:
			// Do some work
			time.Sleep(time.Second * time.Duration(i+1))
			if ctx.Err() != nil {
				// Check if context is canceled before printing
				fmt.Println("Operation canceled.")
				return
			}
			fmt.Printf("Performing operation %v\n", i)
		}
	}
}
func performOperation(ctx context.Context) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * time.Duration(i))
		if ctx.Err() != nil {
			// Check if context is canceled before printing
			fmt.Println("Operation canceled.")
			return
		}
		fmt.Printf("Performing operation %v\n", i)
	}
}

// func UserIDFromContext(ctx context.Context) string {
// 	userID, ok := ctx.Value("userID").(string)
// 	if !ok {
// 		log.Fatal("unable to get userId from context")
// 	}
// 	return userID
// }
