package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age)  // 44
}

func changeMe(z int) { // int is a type
	fmt.Println(z) // passed in the memory address as a value.
	z = 24 // doesn't affect age - would need to pass the memory address to change the value to 24.
}

// When changeMe is called on line 8
// the value 44 is being passed as an argument.