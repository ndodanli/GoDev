package main

import (
	"fmt"
	"time"
)

/*
2 1 0 5 4 6 4
-1 2 -1 5 -1 6 -1

0 1 4 4
2 -1 5 -1 6 -1

3 2 3 1 2 4 5 5 6
2 3 3 1
1 2 2 3 3 4 5 5 6

*/

// func containsDuplicate(nums []int) bool {
// 	n := len(nums)
// 	m := make(map[int]int, n)
// 	for i := 0; i < n; i++ {
// 		m[nums[i]] += 1
// 	}
// 	for _, v := range m {
// 		if v > 1 {
// 			return true
// 		}
// 	}
// 	return false
// }

func containsDuplicate(nums []int) bool {
	n := len(nums)
	var counter int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[i] == nums[j] {
				counter++
			}
		}
		if counter > 1 {
			return true
		}
		counter = 0
	}
	return false
}

func main() {
	start := time.Now()
	fmt.Printf("containsDuplicate([]int{1, 2, 3, 1}): %v\n", containsDuplicate([]int{3, 1}))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
