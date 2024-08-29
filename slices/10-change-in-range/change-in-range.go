package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	for i, v := range x { // длина слайса рассчитывается на первой итерации
		fmt.Printf("%d) %s [l=%d,c=%d]", i, v, len(x), cap(x))
		// fmt.Println(i, v, len(x))
		// fmt.Println(x[i])
		if v == "a" {
			x = append(x[:i], x[i+1:]...)
		}
		fmt.Print("\n")
	}
}
