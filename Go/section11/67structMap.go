package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {"user1", 10},
		2: {"user2", 20},
	}
	fmt.Println(m)

	m2 := map[User]string{
		{"user1", 10}: "Tokyo",
		{"user2", 20}: "LA",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}
}
