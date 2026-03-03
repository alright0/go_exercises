package main

import "fmt"

func fundTwoSum(nums []int, target int) (int, int) {
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := indexMap[complement]; ok {
			fmt.Println(index, i)
			return index, i
		} else {
			indexMap[num] = i
		}
	}
	return 0, 0
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	index1, index2 := fundTwoSum(nums, target)
	fmt.Println(index1, index2)
}
