package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName("B")
	user2.SayName()
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}
