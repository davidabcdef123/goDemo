package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	var num2 []int

	printSlice(num2)
	if (num2 == nil) {
		fmt.Printf("切片是空的")
	}

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v \n", len(x), cap(x), x)
}
