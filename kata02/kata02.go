package main

import "fmt"

// LiniearSearch
// func chop(target int, ls []int) int {
// 	for i, v := range ls {
// 		if target == v {
// 			return i
// 		}
// 	}
// 	return -1
// }

// BinarySearch
func chop(target int, ls []int) int {
	start := 0
	end := len(ls) - 1

	for {
		if start > end {
			return -1
		}
		index := (start + end) / 2
		if target == ls[index] {
			return index
		}

		if ls[index] < target {
			start = index + 1
		} else {
			end = index - 1
		}
	}
}

func main() {
	list := []int{1, 3, 5}

	fmt.Println(chop(3, list))
}
