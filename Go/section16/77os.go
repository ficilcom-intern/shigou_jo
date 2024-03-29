package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//////os.Exit

	os.Exit(1)
	fmt.Println("Start")

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)

	//////os.Open

	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//////os.Args

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("len = %d\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	//////file

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	//////file create

	f2, _ := os.Create("foo.txt")

	f2.Write([]byte("Hello\n"))

	f2.WriteAt([]byte("Golang\n"), 6)

	f2.Seek(0, os.SEEK_END)

	f2.WriteString("Yaah")

	//////file read

	f3, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f3.Close()

	bs := make([]byte, 128)

	n, err := f3.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)
	fmt.Println(string(bs2))

	//////openfile
	f4, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f4.Close()

	bs3 := make([]byte, 128)

	n2, err := f4.Read(bs3)
	fmt.Println(n2)
	fmt.Println(string(bs3))
}
