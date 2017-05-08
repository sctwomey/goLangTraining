package main

import (
	"fmt"
)

func main() { // The idiomatic way to make slices.
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
