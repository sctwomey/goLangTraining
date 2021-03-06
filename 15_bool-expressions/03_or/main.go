package main

import "fmt"

func main() {

	if true || false { // || == Or operator - this is redundant, but is a simple example of its use.
		fmt.Println("This ran")
	}
}

/*
Expressions: produce a value (e.g., an argument in a function call)
- These generally occur horizontally in a program.

Statements: perform an action (e.g., loops and if statements)
- These generally occur vertically in a program.
- A program is basically a sequence of statements.

Also, statements are often made up of expressions, but expressions are not made up of statements.
*/