package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	sl2 := make([]int, 5, 10)

	n := copy(sl2, sl)

	fmt.Println(n, sl2)
}
