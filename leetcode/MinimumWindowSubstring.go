package main

import (
	"fmt"
	"time"
)

// func minWindow(s string, t string) string {
// 	sLen := len(s)
// 	tLen := len(t)
// 	if tLen > sLen {
// 		return ""
// 	}
// 	type TString struct {
// 		MissingChars []byte
// 		StartIndex   int
// 		EndIndex     int
// 	}
// 	arr := make([]TString, 0, sLen-1)
// 	for i := range arr {
// 		arr[i].StartIndex = -1
// 	}
// 	for i := 0; i < sLen; i++ {
// 		for j := 0; j < tLen; j++ {
// 			if s[i] == t[j] {
// 				chars := make([]byte, 0, tLen)
// 				for k := 0; k < tLen; k++ {
// 					// if t[k] != s[i] {
// 					chars = append(chars, t[k])
// 					// }
// 				}
// 				arr = append(arr,
// 					TString{
// 						MissingChars: chars,
// 						StartIndex:   i,
// 					})
// 				for structIndex := len(arr) - 1; structIndex >= 0; structIndex-- {
// 					missingCharsLen := len(arr[structIndex].MissingChars)
// 					if missingCharsLen != 0 {
// 						for l := 0; l < missingCharsLen; l++ {
// 							if arr[structIndex].MissingChars[l] == s[i] {
// 								arr[structIndex].MissingChars[l] = arr[structIndex].MissingChars[missingCharsLen-1]
// 								arr[structIndex].MissingChars = arr[structIndex].MissingChars[:missingCharsLen-1]
// 								missingCharsLen = len(arr[structIndex].MissingChars)
// 								break
// 							}
// 						}
// 						if missingCharsLen == 0 {
// 							if missingCharsLen == tLen {
// 								return s[arr[structIndex].StartIndex : i+1]
// 							} else {
// 								arr[structIndex].EndIndex = i
// 							}
// 						}
// 					}
// 				}
// 				break
// 			}
// 		}
// 	}
// 	minIndexRange := sLen
// 	resultIndex := -1
// 	for i := range arr {
// 		newRange := arr[i].EndIndex - arr[i].StartIndex
// 		if len(arr[i].MissingChars) == 0 && newRange < minIndexRange {
// 			minIndexRange = newRange
// 			resultIndex = i
// 		}
// 	}
// 	if resultIndex != -1 {
// 		return s[arr[resultIndex].StartIndex : arr[resultIndex].EndIndex+1]
// 	}
// 	return ""
// }

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	m := make(map[string]int)
	for i := 0; i < tLen; i++ {
		m[string(t[i])]++
	}

	counter, begin, end, d, head := tLen, 0, 0, sLen+1, 0
	for end < sLen {
		if m[string(s[end])] > 0 {
			counter--
		}
		m[string(s[end])]--
		end++
		for counter == 0 {
			if end-begin < d {
				d = end - begin
				head = begin
				str := s[head : head+d]
				fmt.Printf("str: %v\n", str)
			}
			if m[string(s[begin])] == 0 {
				counter++
			}
			m[string(s[begin])]++
			begin++
		}
	}
	if d == sLen+1 {
		return ""
	} else {
		return s[head : head+d]
	}
}

/*
ADOBECODEBANC
*/
func main() {
	start := time.Now()
	// s, _ := os.ReadFile("test_one.txt")
	// t, _ := os.ReadFile("test_two.txt")
	s := "DAOBECODEBANC"
	//		D AOBEC Y BCA HNC
	//    D A O B E C A B E B  A  N  C
	//    1 2 3 4 5 6 7 8 9 10 11 12 13
	t := "ABC"
	fmt.Printf("Result: %v\n",
		minWindow(string(s), string(t)))
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)
}
