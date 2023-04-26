package main

import "fmt"

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

func main() {
	var t T[int, []*int, *int]
	fmt.Printf("A: %T, B: %T, C: %T\n", t.a, t.b, t.c)
}
