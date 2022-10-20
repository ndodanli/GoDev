package main

import (
	"fmt"
	"time"
)

func moveZeroes(nums []int) {
	n := len(nums)
	selectedIndex := 1
	for i := 0; i < n; i++ {
		if selectedIndex >= n {
			break
		}
		if nums[i] == 0 {
			for j := selectedIndex; j < n; j++ {
				if nums[j] != 0 {
					selectedIndex = j
					break
				}
			}
			nums[i] = nums[selectedIndex]
			nums[selectedIndex] = 0
		}
		selectedIndex++
	}

}
func main() {
	start := time.Now()
	arr := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	moveZeroes(arr)
	fmt.Println(arr)
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
