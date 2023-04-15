package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	// *p = 300
	// fmt.Println(n)

	// n = 200
	// fmt.Println(*p)

	Double2(&n)
	fmt.Println(n)

	Double2(p)
	fmt.Println(*p)

	var sl = []int{1, 2, 3}
	Double3(sl)
	fmt.Println(sl)
}

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}
