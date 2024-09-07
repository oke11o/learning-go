package main

import (
	"fmt"
)

func main() {
	baseSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(baseSlice, len(baseSlice), cap(baseSlice))

	a := baseSlice[1:3:4]
	a[1] = 99
	fmt.Println(baseSlice, a, len(a), cap(a))

	a = append(a, 88)
	fmt.Println(baseSlice, a, len(a), cap(a))

	a = append(a, 77)
	fmt.Println(baseSlice, a, len(a), cap(a))
}
