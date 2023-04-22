package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func main() {
	i := []int{5, 2, 5, 1, 9, 10, 6, 8, 3}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s)

	sort.Ints(i)

	sort.Strings(s)

	fmt.Println(i, s)

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	sort.Slice(el, func(i, j int) bool {
		return el[i].Name < el[j].Name
	})

	fmt.Println("------")
	fmt.Println(el)
	fmt.Println("------")

	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("------")
	fmt.Println(el)
	fmt.Println("------")
}
