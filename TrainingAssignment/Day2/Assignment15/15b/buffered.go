package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("end of program")
	c := make(chan int, 6)
	go Sender(c)
	go Receiver(c)
	time.Sleep(3 * time.Second)

}

func Sender(ch chan int) {

	for i := 1; i < cap(ch); i++ {
		ch <- i
	}

}

func Receiver(ch chan int) {

	for i := 0; i < cap(ch); i++ {
		a := <-ch
		fmt.Println(a)
	}

}
