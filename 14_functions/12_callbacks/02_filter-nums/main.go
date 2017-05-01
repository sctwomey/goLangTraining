package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{} // can be written idiomatically as: var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n) // adding an entry to a slice.
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}