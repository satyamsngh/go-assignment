package main

import (
	"fmt"
	"os"
)

func main() {
	path := `Day2\Assignment8\8a-errorhandling\a.txt`
	err := ErrorHandling(path)
	if err != nil {
		fmt.Println("unable to delete the file , path not found")
		return
	}
	fmt.Println("files are remove")

}

func ErrorHandling(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}

// Create 2 files named a.txt b.txt in your current project directory.
// Create your main.go file.
// Inside main.go file, Write a function that takes the fileName as an input.
// Inside the function, you have to call os.Remove() which would remove a file from the current directory.
// Now call this function from your main function & pass the above created file names to this function.
