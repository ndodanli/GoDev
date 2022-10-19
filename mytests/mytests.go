package main

import (
	"fmt"
)

func calc(h int, citations []int, ch chan int, n int) {
	counter := 0

	for i := 0; i < n; i++ {
		if h <= citations[i] {
			counter++
		}
	}
	if h <= counter {
		ch <- h
	} else {
		ch <- 0
	}
}

func hIndex(citations []int) int {
	n := len(citations)
	ch := make(chan int)
	defer close(ch)
	for i := 0; i < n; i++ {
		go calc(i+1, citations, ch, n)
	}
	var result int = -1
	for i := 0; i < n; i++ {
		v := <-ch
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Printf("hIndex([]int{3, 0, 6, 1, 5}): %v\n", hIndex([]int{0}))
}
