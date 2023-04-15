package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))

	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(sum())

	sl := []int{1, 2, 3}

	fmt.Println(sum(sl...))
}

func sum(s ...int) int {
	n := 0

	for _, v := range s {
		n += v
	}
	return n
}
