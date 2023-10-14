package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go Sender(c, &wg)
	go Receiver(c, &wg)

	wg.Wait()
	fmt.Println("Done")
}

func Sender(ch chan int, wg *sync.WaitGroup) {

	for i := 1; i < 3; i++ {

		ch <- i
		fmt.Println(<-ch)

	}
	wg.Done()
	close(ch)
}

func Receiver(ch chan int, wg *sync.WaitGroup) {

	for i := 1; i < 3; i++ {

		a := <-ch
		fmt.Println(a)

	}
	wg.Done()

}
