package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 10
	fmt.Println(m["k1"])
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	delete(m, "k1")
	_, v3 := m["k1"]
	fmt.Println("v3", v3)
   	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}