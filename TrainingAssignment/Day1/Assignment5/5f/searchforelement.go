package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := os.Args[1:]
	c := []int{}
	for _, v := range a {
		b, _ := strconv.Atoi(v)
		c = append(c, b)

	}
	var s int
	fmt.Scanln(&s)
	search(s, c)

}

func search(s int, a []int) {
	for i, v := range a {
		if v == s {
			fmt.Println("the number was found at", i+1, "position")
			return
		}

	}
	fmt.Println("the number was not found")

}

// Description: Create a program that searches for a specific number in a slice and prints whether it's found or not.
// Instructions:
// Create an integer slice called numbers with at least 5 numbers.
// Choose a number to search for.(take a number input from user using command line)
// Write code to search for that number in the numbers slice.
// Print whether the number was found or not.
