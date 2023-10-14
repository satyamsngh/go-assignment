package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 45, 7, 5}
	slice2 := []int{1, 3, 45, 7, 5}

	c := Comparsion(slice1, slice2)

	fmt.Println(c)

}

func Comparsion(a, b []int) string {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return "slices are not equal"
		}
	}
	return "slices are equal"
}

// Description: Write a program that compares two slices and prints whether they are equal or not.
// Instructions:
// Create two integer slices, slice1 and slice2, with at least 5 numbers in each.
// Write code to compare whether slice1 and slice2 are equal (have the same elements in the same order).
// Print whether the slices are equal or not.
