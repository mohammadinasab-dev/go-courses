package main

import "fmt"

// Component represents the common interface for the original object and decorators.
type Component interface {
	Operation() string
}

// ConcreteComponent is the original object.
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "Doing the core operation."
}

// Decorator adds additional behavior to the original object.
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	// Additional behavior before the core operation.
	// For example, logging or validation.
	result := "Decorator: Pre-operation\n"

	// Call the core operation of the original object.
	result += d.component.Operation()

	// Additional behavior after the core operation.
	// For example, modifying the result or adding extra functionality.
	result += "\nDecorator: Post-operation"

	return result
}

func main() {
	// Create the original object.
	component := &ConcreteComponent{}

	// Decorate the original object with one or more decorators.
	decoratedComponent := &Decorator{component}

	// Call the operation, which includes the additional behavior from the decorator.
	result := decoratedComponent.Operation()
	fmt.Println(result)
}
