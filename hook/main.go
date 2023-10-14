package main

import "fmt"

// Hook function signature
type HookFunc func()

// Function with multiple hooks
func LogPrinter(name string, before HookFunc, after HookFunc) {
	if before != nil {
		before()
	}

	fmt.Println("Executing the main logic")

	if after != nil {
		after()
	}
}

// Custom hook functions
func BeforeHook() {
	fmt.Println("This is the before hook")
}

func AfterHook() {
	fmt.Println("This is the after hook")
}

func main() {

	f := func() {
		fmt.Println("zahra")
	}

	LogPrinter("Hemn", f, AfterHook)

	// Passing nil hooks
	LogPrinter("Hemn", nil, nil)

	
}
