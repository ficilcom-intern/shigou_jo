package main

import "fmt"

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice([]string{"a", "b", "c"})
}
