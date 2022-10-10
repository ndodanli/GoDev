package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	arr := [100000]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	duration := time.Since(start)
	fmt.Println(duration)
	start = time.Now()
	arr1 := [100000]int{}
	for i := 0; i < 100000; i++ {
		arr1[i] = i
	}
	duration = time.Since(start)
	fmt.Println(duration)
}
