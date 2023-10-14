package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 8}

	slice1 = append(slice1, slice2...)

	fmt.Println(slice1)
}

// Description: Create a program that concatenates two slices into a single slice.
// Instructions:
// Create two integer slices, slice1 and slice2, each with at least 3 different numbers.
// Write code to concatenate slice2 to the end of slice1.
// Print the concatenated slice.
