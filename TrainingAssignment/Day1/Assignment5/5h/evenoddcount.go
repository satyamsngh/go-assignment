package main

import "fmt"

func main() {

	number := []int{1, 5, 7, 2, 6, 9, 5, 89, 2}

	odd := 0
	even := 0

	for _, v := range number {
		if v%2 == 0 {
			even += 1

		} else {
			odd += 1
		}
	}

	fmt.Println(odd, even)

}

// Description: Create a program that counts and prints the number of even and odd numbers in a slice.
// Instructions:
// Create an integer slice called numbers with at least 8 numbers (some even and some odd).
// Write code to count the number of even and odd numbers in the numbers slice.
