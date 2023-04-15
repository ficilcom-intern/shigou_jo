package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go receiver(ch1)
	go receiver(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(100 * time.Millisecond)
		i++
	}
	// fmt.Println(<-ch1)
}

func receiver(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}
