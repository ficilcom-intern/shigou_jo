package main

import (
	"Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section18-Todo/app/controllers"
	"Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section18-Todo/app/models"
	_ "Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section18-Todo/config"
	"fmt"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
