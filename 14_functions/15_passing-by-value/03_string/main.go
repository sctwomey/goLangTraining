package main

import "fmt"

func main() {
	name := "Steve"
	fmt.Println(name) // Steve

	changeMe(name)

	fmt.Println(name) // Steve
}

func changeMe(z string) {
	fmt.Println(z) // Steve
	z = "Rocky"
	fmt.Println(z) // Rocky
}