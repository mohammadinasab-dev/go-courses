package main

import "fmt"

/*
	1- slicing an array does not copy the underlying array.
	2- for-range do modification over a copy of respective iteration values.(no changes occured over original entry)
	3- we have three scope when we using for-range statement, outside, for itself, inner.(variable shadowing).
		"for statement block is called an “implicit” block, since it is not surrounded by matching brace brackets { ... }."
	4- for-range declare just one veriable for iteration variable and for each iteration just change the value.
	   "change over each of them does not effect another one".
	5- declare a new variable in each iteration to save the changes for second array assignment.
	6- to change the the original value we should use the slice[index] directly,
	   instead of iterartion variables which is copy of iteration entry.
	7-

*/

func main() {

	demo_01()
	demo_02()
	demo_03()
	demo_04()
	demo_05()
	demo_06()

}

func demo_01() {
	array := [5]int{10, 20, 30, 40, 50}
	slice := array[2:5]

	fmt.Printf("Address of array is %p, value of array is %v\n", &array, array)
	fmt.Printf("Address of slice is %p, value of slice is %v\n", &slice, slice)

	slice[0] = 60
	fmt.Printf("Address of slice is %p, value of slice is %v\n", &slice, slice)
	fmt.Printf("Address of array is %p, value of array is %v\n", &array, array)
}

func demo_02() {
	array := [5]int{10, 20, 30, 40, 50}
	for _, a := range array {
		fmt.Printf("Address of a is: %p, value of a is: %v\n", &a, a)
	}
	fmt.Printf("Address of array is %p, value of array is %v\n", &array, array)
}

func demo_03() {
	a := 1000
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("(func scope): Address of a is: %p, value of a is: %v\n", &a, a)
	for _, a := range array {
		fmt.Printf("(for scope): Address of a is: %p, value of a is: %v\n", &a, a)
		a = a * 10
		fmt.Printf("(inner for scope): Array Address of a is: %p, value of a is: %v\n", &a, a)
	}
	fmt.Printf("Array Address of array is %p, value of array is %v\n", &array, array)
}

func demo_04() {
	array := [5]int{10, 20, 30, 40, 50}
	arrayPtr := []*int{}
	for _, a := range array {
		fmt.Printf("Address of a is: %p, value of a is: %v\n", &a, a)
		arrayPtr = append(arrayPtr, &a)
	}
	fmt.Printf("Address of array is %p, value of array is %v\n", &array, array)
	for _, a := range arrayPtr {
		fmt.Printf("Address of a is: %p, address of a is: %v, value of each index is: %v\n", &a, a, *a)
	}
	fmt.Printf("Address of arrayPtr is %p, value of arrayPtr is %v\n", &arrayPtr, arrayPtr)
}

func demo_05() {
	array := [5]int{10, 20, 30, 40, 50}
	arrayPtr := []*int{}
	for _, a := range array {
		fmt.Printf("Address of a is: %p, value of a is: %v\n", &a, a)
		a := a
		arrayPtr = append(arrayPtr, &a)
	}
	fmt.Printf("Address of array is %p, value of array is %v\n", &array, array)
	for _, a := range arrayPtr {
		fmt.Printf("Address of a is: %p, address of a is: %v, value of each index is: %v\n", &a, a, *a)
	}
	fmt.Printf("Address of arrayPtr is %p, value of arrayPtr is %v\n", &arrayPtr, arrayPtr)
}

func demo_06() {
	array := [5]int{10, 20, 30, 40, 50}
	for i, s := range array {
		array[i] = s * 100
		fmt.Printf("Address of s is: %p, value of s is: %v\n", &s, s)
		fmt.Printf("Address of array[%v] is: %p, value of array[%v] is: %v\n", i, &array[i], i, array[i])

	}
	fmt.Printf("array Address of array is %p, value of array is %v\n", &array, array)
}
