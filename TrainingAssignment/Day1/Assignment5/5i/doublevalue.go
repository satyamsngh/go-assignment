package main

import "fmt"

func main() {
	number := []int{2, 5, 7, 9, 4, 9}

	a := DoubleSlice(number)
	fmt.Println(a)

}

func DoubleSlice(a []int) []int {
	for i, v := range a {
		a[i] = v * v
	}
	return a
}

// Description: Write a program that doubles the values of all elements in a slice.
// Instructions:
// Create an integer slice called numbers with at least 6 different numbers.
// Write code to double the values of all elements in the numbers slice.
// Print the modified slice.
