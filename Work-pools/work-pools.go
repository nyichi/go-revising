package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int){
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const nums = 5
	jobs := make(chan int, nums)
	results := make(chan int, nums)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= nums; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= nums; a++ {
		<-results
	}
}