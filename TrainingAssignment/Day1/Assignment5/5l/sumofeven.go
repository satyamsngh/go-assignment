package main

import "fmt"

func main() {

	number := []int{3, 5, 8, 7, 8, 8}
	sum := 0

	for _, v := range number {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Print(sum)

}

// Description: Write a program that calculates and prints the sum of all even numbers in a slice.
// Instructions:
// Create an integer slice called numbers with at least 6 different numbers (some even and some odd).
// Write code to calculate the sum of all even numbers in the numbers slice.
// Print the sum of even numbers.
