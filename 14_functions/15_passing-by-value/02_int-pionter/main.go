package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // 0x82023c080

	changeMe(&age)

	fmt.Println(&age) // 0x82023c080
	fmt.Println(age)  // 24
}

func changeMe(z *int) { // *int is a type
	fmt.Println(z)  // 0x82023c080 // passed in the memory address as a value.
	fmt.Println(*z) // 44 // * on variable z is an operator and dereferences z to show the literal value.
	*z = 24 // *z says give me the value in z and assign a new value 24.
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 24
}