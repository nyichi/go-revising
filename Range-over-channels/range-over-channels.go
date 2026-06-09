package main

import "fmt"

func main() {
	sentences := make(chan string, 2)
	sentences <- "Hello, World!"
	sentences <- "Welcome to Go channels."
	close(sentences)
	for sentence := range sentences {
		fmt.Println(sentence)
	}
}