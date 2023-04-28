package main

import "fmt"

type Pperson struct {
	Name *string
}

type Person struct {
	Name string
}

func (x Pperson) setName(name string) {
	x.Name = &name
	fmt.Printf("Address is: %p pperson.Name in setName: %v\n", x, *x.Name)

}

func (x *Person) setName(name string) {
	x.Name = name
	fmt.Printf("Address is: %p person.Name in setName: %v\n", x, x.Name)

}

func main() {

	name := "reza"

	fmt.Println("================ pointer to struct ===================")
	person := &Person{
		Name: name,
	}
	// person value:reza   address: 235
	fmt.Printf("Address is: %p person.Name before setName: %v\n", person, person.Name)
	person.setName("sara")
	fmt.Printf("Address is: %p person.Name after setName:: %v\n\n", person, person.Name)

	fmt.Println("================ struct with pointer ===================")
	pperson := Pperson{
		Name: &name,
	}
	fmt.Printf("Address is: %p pperson.Name before setName: %v\n", pperson, *pperson.Name)
	pperson.setName("sara")
	fmt.Printf("Address is: %p pperson.Name after setName: %v\n", pperson, *pperson.Name)

}
