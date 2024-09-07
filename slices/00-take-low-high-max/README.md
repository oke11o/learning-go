# Take low/high/max

Интересует capacity

```go
func main() {
    base := []int{1, 2, 3, 4, 5, 6}
    s := base[1:2:4]
    fmt.Println(s, len(s), cap(s))
}
```