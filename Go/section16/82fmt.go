package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello")

	fmt.Println("Hello!")

	fmt.Println("Hello", "world!")

	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")

	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", "Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("text1.txt")
	defer f.Close()
	fmt.Fprintln(f, "Fprint")
}
