package main

import "fmt"

func reverseArray(arr []int) []int {
	newArr := make([]int, len(arr), len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		newArr[len(arr)-i-1] = arr[i]
	}
	return newArr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	result := reverseArray(arr)

	fmt.Println(result)
}
