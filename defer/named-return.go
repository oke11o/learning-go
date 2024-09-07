package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (res int) {
	res = 8
	defer func() {
		res++
	}()
	return 2
}
