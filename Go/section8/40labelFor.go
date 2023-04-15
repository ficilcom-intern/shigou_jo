package main

import "fmt"

func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("Start")
				break Loop
			}
		}
	}
	fmt.Println("End")

Loop1:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop1
			}

			fmt.Println(i, j, i*j)
		}
		fmt.Println("Not excuted")
	}
}
