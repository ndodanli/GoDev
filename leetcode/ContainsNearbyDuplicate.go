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

//memory friendly
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n && math.Abs(float64(i-j)) <= float64(k); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	if k >= n {
		k = n - 1
	}
	counter := k + 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j <= counter; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
		if counter+1 < n {
			counter++
		}
	}
	return false
}

func main() {
	start := time.Now()
	fmt.Printf("containsNearbyDuplicate([]int{}, 3): %v\n",
		containsNearbyDuplicate([]int{1, 2}, 2))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
