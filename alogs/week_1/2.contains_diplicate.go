package main

import (
	"fmt"
)

func checkDuplicates(nums []int) bool {
	dupMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := dupMap[num]; ok {
			return true
		}
		dupMap[num] = struct{}{}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 7, 8, 9, 8}

	result := checkDuplicates(nums)
	fmt.Println(result)
}
