package main

import "fmt"

func main() {
	numbers := []int{1, 4, 5, 7, 8}
	a := numbers[1]
	for _, v := range numbers {
		if a < v {
			a = v
		}
	}
	fmt.Println("the maximum value in numbers is :", a)

}

// Write a program that finds and prints the maximum element in a slice of integers.
// Instructions:
// Create an integer slice called numbers with at least 5 different numbers in random order..
// Find and print the maximum element in the numbers slice.
