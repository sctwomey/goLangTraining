package main

import (
	"fmt"
)

func main() {
	var student []string // There was nothing that was made underneath, so it's referencing nothing.
	var students [][]string // There was nothing that was made underneath, so it's referencing nothing.
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
