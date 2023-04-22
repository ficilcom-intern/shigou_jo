package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("1st goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("2nd goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("3rd goroutine")
		}
		wg.Done()
	}()

	wg.Wait()
}
