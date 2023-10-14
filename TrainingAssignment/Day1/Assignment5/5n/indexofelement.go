package main

import "fmt"

func main() {
	var v int
	fmt.Scanln(&v)
	number := []int{2, 4, 5, 7, 9, 6}
	val := Element(v, number)
	fmt.Println(val)

}

func Element(s int, a []int) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1

}

// Description: Write a program that finds and prints the index of a specific number in a slice.
// Instructions:
// Create an integer slice called numbers with at least 6 different numbers.
// Choose a number to search for in the numbers slice.
// Write code to find the index of the chosen number. Return -1 if the element is not found.
// Print the index (or a message if the number is not found).
