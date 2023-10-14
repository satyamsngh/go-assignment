package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}
type Person struct {
	Name string
	Address
}

func main() {
	person1 := Person{"Satyam", Address{"23street", "Banglore", "33223"}}
	fmt.Println(person1.Name, person1.City, person1.Street, person1.ZipCode)

}

// Task Description: Create a program that showcases struct composition.

// Instructions:
// 1. Define a struct named Address that has three fields - Street, City, and ZipCode.
// 2. Define another struct named Person that has two fields - Name and embed Address struct in Person.
// 3. Create an instance of the Person struct and initialize both the Name and Address fields in your main function..
// 4. Print the details of the person, including their name and address.
