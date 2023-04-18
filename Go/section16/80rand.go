package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Int())

	fmt.Println(rand.Int31())

	fmt.Println(rand.Int63())

	fmt.Println(rand.Uint32())
}
