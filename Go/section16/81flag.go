package main

import (
	"flag"
	"fmt"
)

func main() {
	//command line option
	//sample
	//go run main.go -n 20 -m message -x

	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "n", 32, "max number")

	flag.StringVar(&msg, "m", "", "message")

	flag.BoolVar(&x, "x", false, "extension")

	flag.Parse()

	fmt.Println("max number = ", max)
	fmt.Println("message = ", msg)
	fmt.Println("extension = ", x)

}
