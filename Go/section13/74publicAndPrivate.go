package main

import (
	"Users/xushao/Desktop/Apps/Gos/shigou_jo/Go/section13/foo"
	"fmt"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return
}

func main() {
	fmt.Println(foo.Max)
	fmt.Println(foo.ReturnMin())

	fmt.Println(appName())

	fmt.Println(Do("AAA"))
}
