package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type T struct {
	// User User
	User
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)
}

func (u *User) SetName() {
	u.Name = "A"
}
