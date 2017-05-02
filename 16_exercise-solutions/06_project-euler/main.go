package main

import "fmt"

func main() {

	var numTrack int // Can be written: numTrack := 0, but this is more idiomatic.

	for i := 0; i < 1000; i++ { // Moves through the range of numbers from 0 to 100.

		if i%3 == 0 { // This is the if statement for checking whether numbers are multiples of 3.
			numTrack += i // Moves through the range of numbers and adds each multiple to the current sum.
		}

		if i%5 == 0 { // This is the if statement for checking whether numbers are multiples of 5.
			numTrack += i // Moves through the range of numbers and adds each multiple to the current sum.
		}
	}

	fmt.Println(numTrack) // Prints the sum of all the multiples of 3 or 5 below 1000.
}

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

// I have to build the .exe file to run most of the programs, so that is why I have a main.exe in the folder.