package main

import (
	"fmt"
)

func main() {
	numbers := []int{}
	numbers = append(numbers, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)
	numbers = append(numbers, 25)
	fmt.Println(numbers)

	fmt.Println(len(numbers), cap(numbers))
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

}

// Description: In this assignment, you'll work with slices and perform basic operations on them.
// Instructions:
// Create an empty integer slice called numbers.
// Add/Append the following numbers to the slice, one at a time, in this order: 5, 10, 15, 20, 25.
// Print the numbers slice after each addition/append.
// Calculate and print the length and capacity of the number slice at the end.
// Remove the number 15 from the numbers slice.(No need to create a new slice)
// Print the number slice after the removal
