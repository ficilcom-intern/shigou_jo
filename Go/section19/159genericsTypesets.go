package main

import "fmt"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	} else {
		return y
	}
}

type MyInt int

func main() {
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max[float64](1.1, 1.3))

	var x, y MyInt = 1, 2
	fmt.Println(Max[MyInt](x, y))
}
