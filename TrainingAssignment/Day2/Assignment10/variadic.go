package main

import "fmt"

func main() {

	total := Sum(1, 3, 2, 4, 56, 8, 6, 88)
	fmt.Println("the total sum is :", total)

}

func Sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// Description: Create a program that defines a function with variadic parameters and uses it.
// Instructions:
// Create a function called sum that takes an arbitrary number of integers and returns their sum.
// Use this function to calculate and print the sum of different sets of numbers.
