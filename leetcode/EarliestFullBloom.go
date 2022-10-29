package main

import (
	"fmt"
	"sort"
	"time"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(growTime)
	indexes := make([]int, n, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}
	var res int
	sort.Slice(indexes, func(i, j int) bool {
		return growTime[indexes[i]] > growTime[indexes[j]]
	})
	var maxDay int = 0
	var daySum int = 0
	for i := 0; i < n; i++ {
		res += plantTime[i]
		daySum += plantTime[indexes[i]]
		if daySum+growTime[indexes[i]] > maxDay {
			maxDay = daySum + growTime[indexes[i]]
		}
	}

	return maxDay
}

/*
ADOBECODEBANC
*/
func main() {
	start := time.Now()
	// s, _ := os.ReadFile("test_one.txt")
	// t, _ := os.ReadFile("test_two.txt")
	// s := "DCCBOAGAACCBABAGGCF"
	//		D AOBEC Y BCA HNC
	//    D A O B E C A B E B  A  N  C
	//    1 2 3 4 5 6 7 8 9 10 11 12 13
	// t := "ABC"
	fmt.Printf("Result: %v\n",
		earliestFullBloom([]int{2, 4}, []int{4, 5}))
	// fmt.Printf("Result: %v\n",
	// 	earliestFullBloom([]int{27, 5, 24, 17, 27, 4, 23, 16, 6, 26, 13, 17, 21, 3, 9, 10, 28, 26, 4, 10, 28, 2}, []int{26, 9, 14, 17, 6, 14, 23, 24, 11, 6, 27, 14, 13, 1, 15, 5, 12, 15, 23, 27, 28, 12}))
	// fmt.Printf("Result: %v\n",
	// 	earliestFullBloom([]int{15, 29, 24, 8, 14, 26, 12, 15, 27, 2, 5, 10, 7, 17, 9, 5, 9, 21, 11, 13, 13, 2, 1, 17, 11}, []int{26, 20, 10, 9, 8, 27, 1, 19, 13, 22, 10, 10, 21, 14, 17, 14, 17, 30, 3, 3, 14, 16, 7, 12, 25}))
	// fmt.Printf("Result: %v\n",
	// 	groupAnagrams([]string{""}))
	// fmt.Printf("Result: %v\n",
	// 	minWindow(string(s), string(t)))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
