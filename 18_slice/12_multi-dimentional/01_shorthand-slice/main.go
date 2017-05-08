package main

import (
	"fmt"
)

func main() {
	student := []string{} // There is an underlying data structure, but there is nothing to reference - it's not nil, but it's not ready.
	students := [][]string{} // There is an underlying data structure, but there is nothing to reference - it's not nil, but it's not ready.
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
