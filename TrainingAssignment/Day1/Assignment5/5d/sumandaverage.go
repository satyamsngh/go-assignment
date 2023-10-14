package main

import "fmt"

func main() {
	initialize := []int{4, 6, 8, 3, 9}
	sum := 0
	for _, v := range initialize {
		sum += v
	}
	fmt.Println(sum)
	a := sum / len(initialize)
	fmt.Println(a)

}

// Create a program that calculates the sum and average of a list of numbers stored in a slice.
// Instructions:
// Create an integer slice called numbers and initialize it with at least 5 numbers.
// Calculate and print the sum of all the numbers in the numbers slice.
// Calculate and print the average of the numbers.
