package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}
	myGreeting["Harleen"] = "Howdy" // Added entry

	fmt.Println(myGreeting)
}
