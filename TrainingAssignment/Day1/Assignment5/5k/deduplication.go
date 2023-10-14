package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{3, 6, 43, 5, 6, 43}
	deduplicated := []int{}
	sort.Ints(numbers)
	deduplicated = append(deduplicated, numbers[0])
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] != numbers[i+1] {
			deduplicated = append(deduplicated, numbers[i+1])
		}
	}
	fmt.Println(deduplicated)

}

// Description: Create a program that removes duplicate elements from a slice.
// Instructions:
// Create an integer slice called numbers with some duplicate values.
// Write code to remove duplicates from the numbers slice.
// Print the deduplicated slice.
