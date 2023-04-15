package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	user1 := User{"user1", 10}
	user2 := User{"user2", 20}
	user3 := User{"user3", 30}
	user4 := User{"user4", 40}

	users := Users{}

	users = append(users, &user1, &user2, &user3, &user4)

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(*u)
	}
}
