package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch1 <- 1
	// ch2 <- "A"
	/*
		v1 := <-ch1
		v2 := <-ch2

		fmt.Println(v1)
		fmt.Println(v2)
	*/

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!!!!")
	default:
		fmt.Println("none")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2

		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
Loop:
	for {
		select {
		case ch3 <- n:
			fmt.Println("case 1")
			n++
		case i3 := <-ch5:
			fmt.Println("received", i3)
		default:
			if n > 10 {
				break Loop
			}
		}

	}
}
