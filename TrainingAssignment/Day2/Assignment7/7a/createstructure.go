package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	employe1 := Employee{12, "sunil", 23, "rajasthan"}
	employe2 := Employee{13, "sandeep", 34, "us"}
	fmt.Println("employe1 details\n", "id", employe1.Id, "\n", "name", employe1.Name, "\n", "age", employe1.Age, "\n", "city", employe1.City)
	fmt.Println("employe2 details\n", "id", employe2.Id, "\n", "name", employe2.Name, "\n", "age", employe2.Age, "\n", "city", employe2.City)

}

// Description: Define a struct for representing an Employee and create instances of it.
// Instructions:
// Define a struct called Employee with fields for Id, Name, Age, and City.
// Create two instances of the Employee struct with different values.
// Print the details of each employee.
