package main

import "fmt"

func main() {

	//Sharing Down --> typically stay in the stack
	x := 5
	inc(&x)
	fmt.Println(x)

	//Sharing Up --> typically escape to the heap
	y := getNumber()
	fmt.Println(*y / 2)

}

func inc(x *int) {
	*x++
}

func getNumber() *int {
	number := 10
	return &number
}


//stack[main...(print)]
//heap[number]