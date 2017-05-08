package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35) // Setting one parameter is both length & capacity. Setting two separates them individually.
	students := make([][]string, 35) // Setting one parameter is both length & capacity. Setting two separates them individually.
	student[0] = "Steve"
	// student = append(student, "Steve")
	fmt.Println(student)
	fmt.Println(students)
}

/*
Setting one parameter is both length & capacity. Setting two separates them individually.
However, if I set two parameters individually (e.g., 20, 100), I will have to use
append to fill in the slice beyond 20 (i.e., at 21, student = append(student, "Steve").
*/