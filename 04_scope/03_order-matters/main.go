package main

import "fmt"

func main() { // inner scope for this program.
	fmt.Println(x) // this is at the block level, order matters for anything in the block level.
	fmt.Println(y) // accessible from the outer scope.

	x := 42
}

var y = 42 // this is at the package level (outer scope for this program), order does not matter for the inner scope.