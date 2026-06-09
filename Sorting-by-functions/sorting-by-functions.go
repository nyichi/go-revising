package main

import (
	"fmt"
	"slices"
	"cmp"
)

func main() {
	fruits := []string{"peach", "apple", "wiki"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println("Sorted fruits:", fruits)
}