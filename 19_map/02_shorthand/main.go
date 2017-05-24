package main

import (
	"fmt"
)

func main() {

	myGreeting := map[string]string { // Composite literal
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}
	fmt.Println(myGreeting["Jenny"])
}
