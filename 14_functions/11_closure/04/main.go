package main

import "fmt"

func wrapper() func() int { // func int is a type - it is returned.
	x := 0 // The more idiomatic way to write this is: var x int to get a 0 value.
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope.

Anonymous function
a function without a name

func expression
assigning a func to a variable
*/