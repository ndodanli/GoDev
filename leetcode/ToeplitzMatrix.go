package main

import (
	"fmt"
	"time"
)

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
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
		isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
