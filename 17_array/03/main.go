package main

import "fmt"

func main() {
	var x [256]int
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 { // Although the array is 256 characters long, the first 50 will only be printed.
			break
		}
	}
}
