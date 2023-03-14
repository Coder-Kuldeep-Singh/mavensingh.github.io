package main

import (
	"fmt"
	"sync"
)

var memo sync.Map

func main() {
	fmt.Println(fib(10))
	fmt.Println(fib(20))
}

func fib(n int) int {
	v, ok := memo.Load(n)
	if ok {
		return v.(int)
	}
	if n < 2 {
		memo.Store(n, n)
		return n
	}

	result := fib(n-1) + fib(n-2)
	memo.Store(n, result)
	return result
}
