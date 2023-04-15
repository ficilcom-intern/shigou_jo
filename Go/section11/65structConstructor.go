package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(*user1)

}

func NewUser(name string, age int) *User {
	return &User{name, age}
}
