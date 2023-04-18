package main

import (
	"fmt"
	"math"
)

func main() {
	//////math const

	fmt.Println(math.Pi)

	fmt.Println(math.Sqrt2)

	fmt.Println(math.MaxFloat32)

	fmt.Println(math.SmallestNonzeroFloat32)

	fmt.Println(math.MaxFloat64)

	fmt.Println(math.SmallestNonzeroFloat64)

	fmt.Println(math.MaxInt64)

	fmt.Println(math.MinInt64)

	//////math functions

	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Cbrt(8))

	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	//////floor ceil
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))

}
