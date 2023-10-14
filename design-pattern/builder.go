package main

// import "fmt"

// type Pizza struct {
// 	dough    string
// 	sauce    string
// 	cheese   string
// 	toppings []string
// }

// type PizzaBuilder struct {
// 	pizza  *Pizza
// 	status int
// }

// func NewPizzaBuilder() *PizzaBuilder {
// 	return &PizzaBuilder{
// 		pizza: &Pizza{},
// 	}
// }

// func (b *PizzaBuilder) SetDough(dough string) *PizzaBuilder {
// 	if b.status != 0 {
// 		panic("status is false")
// 	}
// 	b.pizza.dough = dough
// 	b.status++
// 	return b
// }

// func (b *PizzaBuilder) SetSauce(sauce string) *PizzaBuilder {
// 	b.pizza.sauce = sauce
// 	return b
// }

// func (b *PizzaBuilder) SetCheese(cheese string) *PizzaBuilder {
// 	b.pizza.cheese = cheese
// 	return b
// }

// func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
// 	b.pizza.toppings = append(b.pizza.toppings, topping)
// 	return b
// }

// func (b *PizzaBuilder) Build() *Pizza {
// 	return b.pizza
// }

// func main() {
// 	pizza := NewPizzaBuilder().
// 		SetDough("thin crust").
// 		SetSauce("tomato").
// 		SetCheese("mozzarella").
// 		AddTopping("pepperoni").
// 		AddTopping("mushrooms").
// 		Build()

// 	fmt.Println("Pizza Details:")
// 	fmt.Printf("Dough: %s\n", pizza.dough)
// 	fmt.Printf("Sauce: %s\n", pizza.sauce)
// 	fmt.Printf("Cheese: %s\n", pizza.cheese)
// 	fmt.Println("Toppings:")
// 	for _, topping := range pizza.toppings {
// 		fmt.Println(topping)
// 	}
// }
