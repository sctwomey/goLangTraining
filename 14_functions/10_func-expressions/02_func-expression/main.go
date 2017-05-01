package main

import "fmt"

func main() {

	greeting := func() { // This an anonymous function.
		fmt.Println("Hello world!")
	} // Assigning a variable to an expression is called a func expression.

	greeting()
}
