package main

import "fmt"

func main() {
	studentData := map[string]map[string]string{"sunil": {"age": "34", "grade": "A", "City": "varanasi"}, "sandeep": {"age": "45", "grade": "B", "City": "rajasthan"}, "Aman": {"age": "31", "grade": "C", "City": "varanasi"}}

	for k, v := range studentData {
		fmt.Println(k, v)
	}

}

// Description: Create a program that uses a map of maps to store and retrieve data.
// Instructions:
// Create a map called studentData where each student's name is the key, and the value is another map containing information such as Age, Grade, and City.
// Add at least three students to the studentData map with their respective information.
// Retrieve and print the details of each student.
