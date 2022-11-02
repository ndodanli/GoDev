package main

import (
	"fmt"
	"time"
)

var moves = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func shortestPath(grid [][]int, k int) int {
	current := make([][]int, 0)
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}
	visited[0][0] = k
	current = append(current, []int{0, 0, k})
	next := make([][]int, 0)
	level := 0
	for len(current) != 0 {

		for i := range current {
			if current[i][0] == len(grid)-1 && current[i][1] == len(grid[0])-1 {
				return level
			}
			for _, move := range moves {
				nextx := current[i][0] + move[0]
				nexty := current[i][1] + move[1]
				if nextx >= 0 && nextx < len(grid) && nexty >= 0 && nexty < len(grid[0]) {
					curK := current[i][2]
					if grid[nextx][nexty] == 1 {
						curK--
					}
					if curK >= 0 {
						if visited[nextx][nexty] == -1 || (visited[nextx][nexty] != -1 && curK > visited[nextx][nexty]) {
							visited[nextx][nexty] = curK
							next = append(next, []int{nextx, nexty, curK})
						}
					}
				}
			}
		}
		current, next = next, current
		next = next[:0]
		level++
	}
	return -1
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
		shortestPath([][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, 1))
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
