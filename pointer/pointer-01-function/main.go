package main

import "fmt"

func main() {
	age := 20 // value: 20; address: 123
	pointerToAge := &age // value: 123; address: 289

	fmt.Println("=================================== PassByValue =====================================")
	fmt.Printf("age in main before passByValue; address: %p, value: %d\n", pointerToAge, *pointerToAge)
	passByValue(*pointerToAge)
	fmt.Printf("age in main affter passByValue; address: %p, value: %d\n\n", pointerToAge, *pointerToAge)

	fmt.Println("=================================== PassByPointer =====================================")
	fmt.Printf("age in main before passByPointer; address: %p, value: %d\n", pointerToAge, *pointerToAge)
	passByPointer(pointerToAge)
	fmt.Printf("age in main affter passByPointer; address: %p, value: %d\n\n", pointerToAge, *pointerToAge)

	fmt.Println("=================================== PassByValuePointer =====================================")
	fmt.Printf("age in main before passByValuePointer; address: %p, value: %d\n", pointerToAge, *pointerToAge)
	passByValuePointer(pointerToAge)
	fmt.Printf("age in main affter passByValuePointer; address: %p, value: %d\n\n", pointerToAge, *pointerToAge)

}

func passByValue(age int) {

	fmt.Printf("age in passByValue before assignment; address: %p, value: %d\n", &age, age)
	funcAge := 10
	age = funcAge
	fmt.Printf("age in passByValue affter assignment; address: %p, value: %d\n", &age, age)

}

func passByPointer(age *int) {


	// age || value: 123; address 256
	fmt.Printf("age in passByPointer before assignment; address: %p, value: %d\n", age, *age)
	funcAge := 10
	*age = funcAge
	fmt.Printf("age in passByPointer affter assignment; address: %p, value: %d\n", age, *age)

}

func passByValuePointer(age *int) {

	fmt.Printf("age in passByValuePointer before assignment; address: %p, value: %d\n", age, *age)
	funcAge := 10 //125
	age = &funcAge  
	fmt.Printf("age in passByValuePointer affter assignment; address: %p, value: %d\n", age, *age)

}
