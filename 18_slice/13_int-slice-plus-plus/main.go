package main

import "fmt"

func main() {
	mySlice := make([]int, 1) // I can add to the slice by append.
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++ // Idiomatic way...Says, "Take the value and add 1 to it."
	fmt.Println(mySlice[0])
}
