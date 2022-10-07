// package main

// import (
// 	"fmt"
// 	"sync"

// 	"golang.org/x/tour/tree"
// )

// func Contains[T comparable](s []T, e T) (bool, int) {
// 	for i, a := range s {
// 		if a == e {
// 			return true, i
// 		}
// 	}
// 	return false, -1
// }

// func remove(s []int, i int) []int {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int, wg *sync.WaitGroup, start bool) {
// 	if start {
// 		defer wg.Done()
// 	}
// 	if t == nil {
// 		return
// 	}
// 	if t.Left != nil {
// 		Walk(t.Left, ch, wg, false)
// 	}
// 	ch <- t.Value
// 	if t.Right != nil {
// 		Walk(t.Right, ch, wg, false)
// 		return
// 	}

// }

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {
// 	var wg sync.WaitGroup
// 	ch0, ch1 := make(chan int, 10), make(chan int, 10)
// 	wg.Add(2)
// 	go Walk(t1, ch0, &wg, true)
// 	go Walk(t2, ch1, &wg, true)
// 	wg.Wait()
// 	ch0Arr := make([]int, 10)
// 	ch1Arr := make([]int, 10)
// 	for i := 0; i < 10; i++ {
// 		ch0Arr[i] = <-ch0
// 		ch1Arr[i] = <-ch1
// 	}
// 	for _, v := range ch0Arr {
// 		if res, index := Contains(ch1Arr, v); res {
// 			ch1Arr = remove(ch1Arr, index)
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	// ch0, ch1 := make(chan int, 10), make(chan int, 10)
// 	t1 := tree.New(2)
// 	t2 := tree.New(2)
// 	t1.Value = 1
// 	// go Walk(t1, ch0)
// 	// go Walk(t2, ch1)

// 	fmt.Println(Same(t1, t2))

// }
