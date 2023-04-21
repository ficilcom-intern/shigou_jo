package main

import (
	"Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section18-Todo/app/models"
	_ "Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section18-Todo/config"
	"fmt"
)

func main() {
	// fmt.Println(config.Config)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)

	fmt.Println(u)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)

	fmt.Println(u)
}
