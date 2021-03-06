package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// If the ... is after the element, then it is an argument.
// If the ... is before the element, then it is a parameter.
// This is an alternative way to pass arguments. This happens to be from a slice.