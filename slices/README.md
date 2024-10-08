# Tasks



- 0. [Take low/high/max](00-take-low-high-max/README.md)

## 1
https://dev.to/crusty0gphr/tricky-golang-interview-questions-part-1-slice-header-3oo0

```go
package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	nums := make([]int, 0, 2)
	fmt.Println(nums, len(nums), cap(nums))

	appendSlice(nums, 1024)
	fmt.Println(nums, len(nums), cap(nums))

	mutateSlice(nums, 0, 512)
	fmt.Println(nums, len(nums), cap(nums))
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
	fmt.Println(sl, len(sl), cap(sl))
}

func mutateSlice(sl []int, i, val int) {
	sl[i] = val
} 
```

## 2 
https://parsiya.net/blog/2020-05-17-go-slices-and-their-oddities/

```go
package main

import "fmt"

func printSlice(s string, a []int) {
    fmt.Printf("%p - %v\tlen:%d\tcap:%d\t%s\n", a, a, len(a), cap(a), s)
}

func surprise(a []int) {
    printSlice("Inside surprise, before assignment", a)
    for i := range(a) {
        a[i] = 5
    }
    printSlice("Inside surprise, after assignment", a)
}
// Quiz #1
func main() {
    a := []int{1, 2, 3, 4}
    printSlice("Inside main, before surprise", a)
    surprise(a)
    printSlice("Inside main, after surprise", a)
}
```

## 3
https://parsiya.net/blog/2020-05-17-go-slices-and-their-oddities/

```go
package main

import "fmt"

func printSlice(s string, a []int) {
	fmt.Printf("%p - %v\tlen:%d\tcap:%d\t%s\n", a, a, len(a), cap(a), s)
}

func surprise(a []int) {
	printSlice("Inside surprise, before append", a)
	a = append(a, 5)
	printSlice("Inside surprise, after append", a)
	printSlice("Inside surprise, before assignment", a)
	for i := range a {
		a[i] = 5
	}
	printSlice("Inside surprise, after assignment", a)
}

// Quiz #2
func main() {
	a := []int{1, 2, 3, 4}
	printSlice("Inside main, before surprise", a)
	surprise(a)
	printSlice("Inside main, after surprise", a)
}
```

## 4
https://parsiya.net/blog/2020-05-17-go-slices-and-their-oddities/

```go
package main

import "fmt"

func printSlice(s string, a []int) {
	fmt.Printf("%p - %v\tlen:%d\tcap:%d\t%s\n", a, a, len(a), cap(a), s)
}

func surprise(a []int) {
	printSlice("Inside surprise, before append", a)
	a = append(a, 5)
	printSlice("Inside surprise, after append", a)
	printSlice("Inside surprise, before assignment", a)
	for i := range a {
		a[i] = 5
	}
	printSlice("Inside surprise, after assignment", a)
}

// Quiz #3
func main() {
	a := []int{1, 2, 3, 4}
	printSlice("Inside main, before surprise", a)
	surprise(a)
	printSlice("Inside main, after surprise", a)
	printSlice("Inside main, before append", a)
	a = append(a, 5)
	printSlice("Inside main, after append", a)
}
```

## 5
https://parsiya.net/blog/2020-05-17-go-slices-and-their-oddities/

```go
package main

import "fmt"

func printSlice(s string, a []int) {
	fmt.Printf("%p - %v\tlen:%d\tcap:%d\t%s\n", a, a, len(a), cap(a), s)
}

func surprise(a []int) {
	printSlice("Inside surprise, before append", a)
	a = append(a, 5)
	printSlice("Inside surprise, after append", a)
	printSlice("Inside surprise, before assignment", a)
	for i := range a {
		a[i] = 5
	}
	printSlice("Inside surprise, after assignment", a)
}

// Quiz #4
func main() {
	a := []int{1, 2, 3, 4}
	printSlice("Inside main, before append", a)
	a = append(a,5)
	printSlice("Inside main, after append", a)
	printSlice("Inside main, before surprise", a)
	surprise(a)
	printSlice("Inside main, after surprise", a)
}
```


## 6
https://medium.com/@ninucium/go-interview-questions-part-2-slices-87f5289fb7eb

```go
package main

import "fmt"

func main() {
  x := []int{1, 2, 3, 4, 5}
  x = append(x, 6)
  x = append(x, 7)
  a := x[4:]
  y := alterSlice(a)
  fmt.Println(x)
  fmt.Println(y)
}

func alterSlice(a []int) []int {
  a[0] = 10
  a = append(a, 11)
  return a
}
```

## 7

Как преобразовать []myStruct в []any?

```go
package main

import "fmt"

type myStruct struct {
	x int
}

func printSlice(v []any) {
	fmt.Println(v)
}

func main () {
	list := []myStruct{{1},{2}}

	printSlice(list)
}
```

## 8

Реализуй стек на слайсах
```go
package stack

import "fmt"

type Stack []int

func NewStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) Push(data int) {
	*s = append(*s, data)
}

func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, nil
}
```

- [get-slice-from-other-capacity](09-get-slice-from-other-capacity/README.md)
- [change-in-range](./10-change-in-range/README.md)
