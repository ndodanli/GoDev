package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	var first, second uint32 = 0, 1
	if n == 0 {
		return 0
	}
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return int(second)
}

func main() {
	start := time.Now()
	fmt.Printf("Result: %v\n",
		fib(30))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
