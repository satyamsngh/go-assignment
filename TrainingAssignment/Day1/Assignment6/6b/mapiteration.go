package main

import "fmt"

func main() {
	fruits := map[string]int{"apple": 2, "banana": 5, "grapes": 4, "orange": 3}

	for key, val := range fruits {
		
			fmt.Println(key,val)
		

	}
}

// Description: Write a program that iterates over a map and prints its key-value pairs.
// Instructions:
// Create a map called fruits with fruit names as keys and their corresponding quantities as values.
// Add at least four different fruits and quantities to the map.
// Write code to iterate over the fruits map and print each fruit and its quantity.
