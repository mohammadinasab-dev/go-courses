package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()

		// Outer goroutine
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			fmt.Println("Starting outer goroutine")
			time.Sleep(1 * time.Second)
			panic("Oops! Panic occurred in outer goroutine")
		}()

		time.Sleep(2 * time.Second)
		fmt.Println("This line will be executed in the inner goroutine")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine finished")
}
