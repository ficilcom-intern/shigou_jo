package main

import "fmt"

type Vector[T any] []T

type IntVector Vector[int]

func main() {
	var v Vector[int] = Vector[int]{10, 20, 30}
	fmt.Println(v)

	v2 := Vector[float64]{1.3, 3.4, 4.5}
	fmt.Println(v2)

	v3 := IntVector{1, 2, 3}
	fmt.Println(v3)
}
