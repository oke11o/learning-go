package main

import (
	"context"
	"fmt"
)

func main() {
	//var ch chan int
	//go func() {
	//	fmt.Println("goroutine 1")
	//	ch <- 1
	//}()
	//go func() {
	//	fmt.Println("goroutine 2")
	//	v := <-ch
	//	fmt.Println(v)
	//}()

	//var ch chan int
	//close(ch)
	//ch <- 1

	//go func() {
	//	fmt.Println("goroutine")
	//	var ch chan int
	//	ch <- 1
	//}()

	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel
	//cancel()
	input := make(chan int, 1)
	close(input)
	//input <- 8

	f1(ctx, input)
	f2(ctx, input)

	//time.Sleep(5 * time.Second)
	fmt.Println("DONE")
}

func f1(ctx context.Context, input <-chan int) {
	select {
	case <-ctx.Done():
		fmt.Println("Ctx done 1")
	case v := <-input:
		fmt.Println(v)
	default:
	}
}

func f2(ctx context.Context, input <-chan int) {
	select {
	case <-ctx.Done():
		fmt.Println("Ctx done 2")
	case v := <-input:
		fmt.Println(v)
	default:
	}
}
