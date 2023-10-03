package main

import (
	"context"
	"fmt"
	"time"
)

/*

	- A context is propagated or cascaded cancellation signals, deadlines, and other context-related values 
	  down to child contexts, not up

*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second/2)
	defer cancel()

	//redis as db
	ctxDB, cancelCtxDB := context.WithTimeout(ctx, time.Second*1)
	defer cancelCtxDB()

	ctxIO, cancelCtxIO := context.WithTimeout(ctx, time.Second*3)
	defer cancelCtxIO()

	go performDBOperationGraceFully(ctxDB)

	go performIOOperationGraceFully(ctxIO)

	func() {
		time.Sleep(time.Second * 4)
	}()

	fmt.Println("Main goroutine terminated.")
}

func performDBOperationGraceFully(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("DB Operation canceled.")
		return
	case <-time.NewTimer(time.Second * 1).C:
		fmt.Println("Performing DB operation")
	}
}

func performIOOperationGraceFully(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("IO Operation canceled.")
		return
	case <-time.NewTimer(time.Second * 5).C:
		fmt.Println("Performing IO operation")
	}
}
