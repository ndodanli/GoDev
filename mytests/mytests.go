package main

import (
	"fmt"
	"log"
)

type customString string
type customInt int

func (x customString) String() string {
	return "dsadsad"
}

func Println[T ~int | ~string](v T) {
	log.Println(v)
}

func printMemoryAddress[T any](v T) {
	log.Printf("memory address: %p", v)
}
func main() {
	// arr0 := [2]int{1, 2}
	slice0 := make([]int, 10, 15)
	fmt.Printf("value: %v length: %v cap: %v\n", slice0, len(slice0), cap(slice0))
	slice0Ptr := slice0
	printMemoryAddress(slice0)
	printMemoryAddress(slice0Ptr)
	slice0[0] = 1
	slice0 = append(slice0, 1, 2, 3, 4, 5)
	fmt.Printf("value: %v length: %v cap: %v\n", slice0, len(slice0), cap(slice0))
	printMemoryAddress(slice0)
	printMemoryAddress(slice0Ptr)
	fmt.Println(slice0Ptr)
	slice0 = append(slice0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Printf("value: %v length: %v cap: %v\n", slice0, len(slice0), cap(slice0))
	printMemoryAddress(slice0)
	printMemoryAddress(slice0Ptr)
	fmt.Println(slice0Ptr)
	// fmt.Printf("arr0: %p\n", arr0)
	// fmt.Printf("slice0: %p\n", slice0)
}
