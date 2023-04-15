package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 2)

	/*
		ch1 <- 1

		close(ch1)

		// ch1 <- 1
		// fmt.Println(<-ch1)

		i, ok := <-ch1
		fmt.Println(i, ok)

		i2, ok2 := <-ch1
		fmt.Println(i2, ok2)
	*/

	go receiver("1.goroutine", ch1)
	go receiver("2.goroutine", ch1)
	go receiver("3.goroutine", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "end")
}
