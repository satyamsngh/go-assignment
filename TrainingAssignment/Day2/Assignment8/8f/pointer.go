package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	student1 := Student{"satyam", 55}
	UpdatingAge(&student1)
	fmt.Println(student1)

}

func UpdatingAge(s *Student) {
	s.Age += 1

}

// Define a struct called Student with fields for Name (string) and Age (int).
// Write a function that takes a pointer to a Student as an argument and updates the Age field by incrementing it by 1.
// In your main program, create an instance of the Student struct, pass it by reference to the function, and then print the updated Age of the student.
