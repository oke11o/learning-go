package main

import (
	"fmt"
)

func main() {
	baseSlice := []int{1, 2, 3, 4}
	fmt.Println(baseSlice, len(baseSlice), cap(baseSlice))
	a0 := append(baseSlice, 5, 6, 7, 8)
	fmt.Println(a0, len(a0), cap(a0))
	a1 := append(baseSlice, 5, 6, 7, 8, 9)
	fmt.Println(a1, len(a1), cap(a1))
}

func main2() {
	baseSlice := []int{1, 2, 3}
	fmt.Println(baseSlice, len(baseSlice), cap(baseSlice))

	a0 := append(baseSlice, 5, 6, 7)
	fmt.Println(a0, len(a0), cap(a0))

	a1 := append(baseSlice, 5, 6, 7, 8)
	fmt.Println(a1, len(a1), cap(a1))

	a2 := append(baseSlice, 5, 6, 7, 8, 9)
	fmt.Println(a2, len(a2), cap(a2))

	a3 := append(baseSlice, 5, 6, 7, 8, 9, 0)
	fmt.Println(a3, len(a3), cap(a3))

	a4 := append(baseSlice, 5, 6, 7, 8, 9, 0, 1)
	fmt.Println(a4, len(a4), cap(a4))
}
