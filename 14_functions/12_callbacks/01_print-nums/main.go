package main

import "fmt"

func visit(numbers []int, callback func(int)) { // visit is the name of the function that takes two parameters (numbers and callback).
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing a func as an argument.
// This is functional programming, but is not ideal for running this type of program in Go.