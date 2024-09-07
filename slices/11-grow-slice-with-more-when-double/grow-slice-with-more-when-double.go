package main

import (
	"fmt"
)

func main() {
	baseSlice := []int{1, 2, 3}
	fmt.Println(baseSlice, len(baseSlice), cap(baseSlice))

	a0 := append(baseSlice, 5, 6, 7)
	fmt.Println(a0, len(a0), cap(a0))

	a1 := append(baseSlice, 5, 6, 7, 8)
	fmt.Println(a1, len(a1), cap(a1))
}

func main2() {
	baseSlice := []int{1, 2, 3, 4}
	fmt.Println(baseSlice, len(baseSlice), cap(baseSlice))
	a0 := append(baseSlice, 5, 6, 7, 8)
	fmt.Println(a0, len(a0), cap(a0))
	a1 := append(baseSlice, 5, 6, 7, 8, 9)
	fmt.Println(a1, len(a1), cap(a1))
}
