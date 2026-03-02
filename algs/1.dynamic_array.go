package main

import "fmt"

type DynamicArray struct {
	data []int
}

func push(arr *DynamicArray, value int) {
	arrSize := len(arr.data)
	if arrSize == cap(arr.data) {
		arrCap := max(cap(arr.data)*2, 1)

		newArr := make([]int, arrSize, arrCap)
		copy(newArr, arr.data)
		arr.data = newArr
	}

	arr.data = arr.data[:arrSize+1]
	arr.data[arrSize] = value
}

func main() {
	dynamicArray := DynamicArray{
		data: make([]int, 0, 0),
	}

	push(&dynamicArray, 1)
	push(&dynamicArray, 2)
	push(&dynamicArray, 3)
	push(&dynamicArray, 4)
	push(&dynamicArray, 5)

	fmt.Println(dynamicArray, cap(dynamicArray.data))
}
