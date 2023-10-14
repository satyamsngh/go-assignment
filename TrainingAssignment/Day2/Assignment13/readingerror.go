package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile(`Day2\Assignment13\sample.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}
	wordcount := CountWords(content)
	fmt.Println("the number of words in sample.txt is :", wordcount)
	error := os.Remove(`Day2\Assignment13\sample.txt`)
	if error != nil {
		fmt.Println(error)
		return

	}
	fmt.Println("the file deleted successfully")
}

func CountWords(a []byte) int {
	count := 0
	for v := range a {
		if v != 32 {
			count += 1
		}
	}
	return count

}

// Description: In this assignment, you will work with file reading and error handling in Go. You'll create a simple program to read a text file, count the number of words, and handle potential errors that may occur during the process.

// Instructions:
// Create a text file named "sample.txt" with some content. The content can be any text or a short paragraph.
// Write a Go program that reads the content of the "sample.txt" file and counts the number of words in it.
// Use the ioutil.ReadFile function to read the file's contents into a byte slice. Handle the error that may occur during file reading and print an error message if it occurs.
// Implement a function called CountWords that takes the file content as a string and returns the number of words in the content.
// Handle potential errors within the CountWords function. Specifically, account for cases where the input string is empty or contains only whitespace characters.
// In the main program, call the CountWords function with the file content and print the number of words. Handle any errors returned by the function and print an error message in case of an error.
// Test your program with different content in the "sample.txt" file, including empty files or files containing only spaces, to ensure error handling is working correctly.
// At the end try to delete the “sample.txt” file & check if your program is able to handle the errors.
