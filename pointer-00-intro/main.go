package main

import "fmt"

func main() {
	var (
		Age int  = 10
		pToAge *int = &Age
	)

	Age = 10
	age2 :=20
	pToAge = &age2

	fmt.Printf("a is %d\n", Age)
	fmt.Printf("p is %d\n", *pToAge)

}
