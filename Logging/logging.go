package main

import (
	// "fmt"
	"log"
)

func main() {
	log.Println("Hello, World!")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	log.Println("with micro")
}