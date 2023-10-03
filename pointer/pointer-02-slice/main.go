package main

import "fmt"

func main() {

	fmt.Println("========================================================================")
	firstSlice := []int{10, 20, 30, 40, 50}
	fmt.Printf("firstSlice: %v, len: %d, cap: %d\n\n", firstSlice, len(firstSlice), cap(firstSlice))

	fmt.Println("=================================== Create second slice by first =====================================")
	secondSlice := firstSlice[:2]
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("firstSlice: %v, len: %d, cap: %d\n\n", firstSlice, len(firstSlice), cap(firstSlice))

	fmt.Println("=================================== Change a single value of second slice =====================================")
	secondSlice[0] = 90
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("firstSlice: %v, len: %d, cap: %d\n\n", firstSlice, len(firstSlice), cap(firstSlice))

	fmt.Println("=================================== Append some values into the second slice less than capacity =====================================")
	secondSlice = append(secondSlice, 100, 110, 120)
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("firstSlice: %v, len: %d, cap: %d\n\n", firstSlice, len(firstSlice), cap(firstSlice))

	fmt.Println("=================================== Append some values into second slice more than capacity =====================================")
	secondSlice = append(secondSlice, 100, 110, 120)
	secondSlice[0] = 9000
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("firstSlice: %v, len: %d, cap: %d\n\n", firstSlice, len(firstSlice), cap(firstSlice))

	fmt.Println("=================================== Assign third slice to the second =====================================")
	thirdSlice := []int{60, 70, 80}
	secondSlice = thirdSlice
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("firstSlice: %v, len: %d, cap: %d\n\n", firstSlice, len(firstSlice), cap(firstSlice))

	changeSliceFunc(secondSlice)
	fmt.Println("=================================== OutSide the changeSlice function =====================================")
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n\n", secondSlice, len(secondSlice), cap(secondSlice))

	pointerToSliceFunc(&secondSlice)
	fmt.Println("=================================== OutSide the pointerSlice function =====================================")
	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n\n", secondSlice, len(secondSlice), cap(secondSlice))

}

func changeSliceFunc(changeSlice []int) {
	

	fmt.Println("=================================== Inside the changeSlice function =====================================")
	fmt.Printf("changeSlice: %v, len: %d, cap: %d\n", changeSlice, len(changeSlice), cap(changeSlice))
	changeSlice[0] = 130
	fmt.Printf("changeSlice: %v, len: %d, cap: %d\n\n", changeSlice, len(changeSlice), cap(changeSlice))

}

func pointerToSliceFunc(pointerSlice *[]int) {
	// pointerSlice || value:123 Address:203
	

	
	fmt.Println("=================================== Inside the pointerToSliceFunc function =====================================")
	fmt.Printf("pointerSlice: %v, len: %d, cap: %d\n", *pointerSlice, len(*pointerSlice), cap(*pointerSlice))
	//*[203] 
	*pointerSlice = []int{600, 500}
	fmt.Printf("pointerSlice: %v, len: %d, cap: %d\n\n", *pointerSlice, len(*pointerSlice), cap(*pointerSlice))
}

func CopySlice(slice []int) []int {

	var temp []int
	for i, s := range slice {
		temp[i] = s
	}
	return temp
}

// type slice struct {
// Length   int
// Capacity int
// Header   *byte
// pointer *referenceAddress
// }
