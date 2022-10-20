package main

import (
	"fmt"
	"time"
)

func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	for i := 0; i < hLen; i++ {
		if haystack[i] == needle[0] {
			exist := i
			nCounter := 1
			if hLen-i >= nLen {
				for j := i + 1; j < i+nLen; j++ {
					if haystack[j] != needle[nCounter] {
						exist = -1
					}
					nCounter++
				}
			} else {
				exist = -1
			}
			if exist != -1 {
				return exist
			}
		}
	}
	return -1
}

func main() {
	start := time.Now()
	fmt.Printf("strStr(\"sadbutsad\", \"sad\"): %v\n", strStr("3f2f3gsdf", "d"))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
