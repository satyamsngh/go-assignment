package main

import "fmt"

func main() {
	studentGrades := map[string]float64{"Anil": 5.7, "Sunil": 7.5, "Shivam": 6.7}

	for _, v := range studentGrades {
		fmt.Printf("the students grades are %v", v)
		fmt.Println()

	}

	delete(studentGrades, "Anil")
	fmt.Println("after deleting the anil ",studentGrades)

}

// Description: Create a simple program that uses a map to store and retrieve information.
// Instructions:
// Create a map called studentGrades with student names as keys and their corresponding grades as values.
// Add at least three students to the map with their grades.
// Print the grades of each student.
// Delete one student's entry from the map.
// Now print the students map again.
