package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	s := array[1:2:6]
	fmt.Println(s, "lencap =", len(s), cap(s))
	a := s[1:5]
	fmt.Println(a, "lencap =", len(a), cap(a))
}
