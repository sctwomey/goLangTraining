package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string) // Also, can be stated as a composite literal: var myGreeting = map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}

// Composite literal == {} appended to the end of the var myGreeting = map[string]string...