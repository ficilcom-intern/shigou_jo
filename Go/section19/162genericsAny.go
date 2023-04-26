package main

import "fmt"

func main() {
	var a any = 1
	a = "a"
	a = true
	fmt.Println(a)
}
