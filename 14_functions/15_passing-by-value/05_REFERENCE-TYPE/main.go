package main

import "fmt"

func main() {
	m := make([]string, 1, 25) // slice, length, capacity
	fmt.Println(m) // [ ] // slice
	changeMe(m) // passing the slice into function changeMe.
	fmt.Println(m) // [Steve]
}

func changeMe(z []string) { // receives the slice ([]slice), then assigns the slice to variable z.
	z[0] = "Steve" // this is saying, "the item in position 0, make it Steve.
	fmt.Println(z) // [Steve]
}

/*
Functions
In Go, everything is passed by value. The type is what is important. Is it a reference type or not a reference type.
Not a reference type: int, string, bool, etc. are primitive values. So, if you want to change the value for
anything that is not a reference type, you need to pass the memory address.

If the element is a reference type (e.g., slice), then there is no need to pass the memory address.

Slice is pointer to an underlining array (a reference type pointing to an underlying array).
*/