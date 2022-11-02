package main

import (
	"fmt"
	"sort"
	"time"
)

func reorganizeString(s string) string {
	m := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		m[s[i]]++
	}

	var max int
	var sum int
	for _, el := range m {
		if el > max {
			max = el
		}
		sum += el
	}

	if sum-max < max-1 {
		return ""
	}

	mLen := len(m)
	keys := make([]byte, 0, mLen)
	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	res := make([]byte, 0, n+n)

	for i := 0; i < m[keys[0]]; i++ {
		res = append(res, keys[0])
		for j := 1; j < mLen; j++ {
			if m[keys[j]] > 0 {
				res = append(res, keys[j])
				res = append(res, byte(0))
				m[keys[j]]--
			}
		}
	}
	for i, j := len(res)-1, 2; i > 0; i, j = i-1, j+2 {
		if res[i] == res[i-1] {
			if res[j+1] == res[i] {
				res[j+3], res[i] = res[i], res[j+3]
				j += 3
			} else {
				res[j], res[i] = res[i], res[j]
			}
		} else {
			break
		}
	}

	res2 := make([]byte, 0, n)
	for i := 0; i < len(res); i++ {
		if res[i] != byte(0) {
			res2 = append(res2, res[i])
		}
	}

	return string(res2)
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
		reorganizeString("bbbbbxbwhbbbbmsybtbbbbbkncyybnjbhxbbrxibcjybb"))
	// fmt.Printf("len(reorganizeString(\"vvvlo\")): %v\n", len(reorganizeString("vvvlo")))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
