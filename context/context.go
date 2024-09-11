package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx2 := context.WithValue(ctx, 0, "ctx2")
	ctxTr := context.WithValue(ctx2, 0, "SpanTrace")
	ctx4 := context.WithValue(ctxTr, 0, "ctx4")

	fmt.Println(ctx.Value(0))
	fmt.Println(ctx2.Value(0))
	fmt.Println(ctxTr.Value(0))
	fmt.Println(ctx4.Value(0))
}
