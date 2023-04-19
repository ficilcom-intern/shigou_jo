package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("start")
	time.Sleep(time.Second)
	fmt.Println("end")
	ch <- "result"
}

func main() {
	ch := make(chan string)

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("###Error###")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}

	}

	fmt.Println("out of loop")
}
