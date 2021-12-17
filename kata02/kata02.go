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

// BinarySearch v1
// func chop(target int, ls []int) int {
// 	start := 0
// 	end := len(ls) - 1
//
// 	for {
// 		if start > end {
// 			return -1
// 		}
// 		index := (start + end) / 2
// 		if target == ls[index] {
// 			return index
// 		}
//
// 		if ls[index] < target {
// 			start = index + 1
// 		} else {
// 			end = index - 1
// 		}
// 	}
// }

//BinarySearch v2
func chop(target int, ls []int) (result int) {
	mediam := len(ls) / 2

	switch {
	case len(ls) == 0:
		result = -1
	case ls[mediam] < target:
		result = chop(target, ls[mediam+1:])
		if result == -1 {
			return -1
		}
		result += mediam + 1
	case ls[mediam] > target:
		result += mediam
		result = chop(target, ls[:mediam])
		if result == -1 {
			return -1
		}
	default:
		result = mediam
	}
	return
}

func main() {
	list := []int{3}

	fmt.Println(chop(5, list))
}
