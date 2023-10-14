package main

// import "fmt"

// // Strategy interface
// type PaymentStrategy interface {
// 	Pay(amount float64) error
// }

// // Concrete strategies
// type CreditCardPaymentStrategy struct{}

// func (ccps *CreditCardPaymentStrategy) Pay(amount float64) error {
// 	fmt.Printf("Paid $%.2f using credit card.\n", amount)
// 	return nil
// }

// type PayPalPaymentStrategy struct{}

// func (ppps *PayPalPaymentStrategy) Pay(amount float64) error {
// 	fmt.Printf("Paid $%.2f using PayPal.\n", amount)
// 	return nil
// }

// // Context
// type PaymentContext struct {
// 	strategy PaymentStrategy
// }

// func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
// 	pc.strategy = strategy
// }

// func (pc *PaymentContext) Pay(amount float64) error {
// 	return pc.strategy.Pay(amount)
// }

// func main() {
// 	// Create a payment context with the default strategy
// 	context := &PaymentContext{strategy: &CreditCardPaymentStrategy{}}

// 	// Use the context to pay an amount
// 	amount := 100.0
// 	err := context.Pay(amount)
// 	if err != nil {
// 		fmt.Printf("Payment failed: %v\n", err)
// 	}

// 	// Change the strategy at runtime
// 	context.SetStrategy(&PayPalPaymentStrategy{})
// 	err = context.Pay(amount)
// 	if err != nil {
// 		fmt.Printf("Payment failed: %v\n", err)
// 	}
// }
