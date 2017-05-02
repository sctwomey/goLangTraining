package main

import "fmt"

func main() {
	greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(greatest)
}

func max(numbers ...int) int {
	//fmt.Printf("%T\n", numbers) // checking the type of numbers.
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}