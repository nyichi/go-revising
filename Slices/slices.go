package main

import (
	"fmt"
)

func main(){
	var s []string
	fmt.Println(s == nil)

	s = make([]string, 3)
	fmt.Println(len(s))

		
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

	l := s[2:5]
   	fmt.Println("sl1:", l)
}