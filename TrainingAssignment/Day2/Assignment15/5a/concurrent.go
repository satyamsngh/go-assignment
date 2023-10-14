package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("end of program")
	c := make(chan int)
	go Sender(c)
	go Receiver(c)
	time.Sleep(3 * time.Second)

}

func Sender(ch chan int) {

	for i := 1; i < 6; i++ {
		ch <- i
	}

}

func Receiver(ch chan int) {

	for i := 0; i < 5; i++ {
		a := <-ch
		fmt.Println(a)
	}

}

// Create a Go program with two functions: sender and receiver.
// The sender function should use a goroutine to send a sequence of numbers (e.g., 1 to 5) to a channel.
// The receiver function should use a goroutine to receive numbers from the channel and print each received number.
// Use an unbuffered channel to ensure that the sender and receiver synchronize their operations.
// In the main function, start both the sender and receiver goroutines.
// Use the time.Sleep function to allow the goroutines to execute. You can sleep for a few seconds to ensure the sender and receiver have time to complete.(Using time.Sleep in your code for synchronization is a bad practice but to keep this example simple you can use time.Sleep here for now. We will remove it in the next step)
// Print a message in the main function to indicate the end of the program.
