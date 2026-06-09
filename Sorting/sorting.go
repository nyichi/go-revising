package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{5, 2, 9, 1, 5, 6}
	slices.Sort(numbers)
	fmt.Println("Sorted numbers:", numbers)
}