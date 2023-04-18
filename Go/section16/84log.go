package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	log.Fatal("Log\n")
	log.Fatalln("Log2")
	log.Fatalf("Log%d\n", 3)

	log.Panic("Log\n")
	log.Panicln("Log2\n")
	log.Panicf("Log%d\n", 3)

	f, err := os.Create("test.log")
	if err != nil {
		return
	}

	log.SetOutput(f)
	log.Println("Write into file")

	log.SetOutput(os.Stdout)

	log.SetFlags(log.LstdFlags)
	log.Println("A")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG]")
	log.Println("D")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")
	log.Println("message")

	_, err1 := os.Open("d")
	if err != nil {
		// log.Fatalln("Exit", err1)
		logger.Fatalln("Exit", err1)
	}

}
