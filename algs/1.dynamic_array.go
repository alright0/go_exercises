package main

import (
	"fmt"
)

type DArray struct {
	data []int
}

func push(arr *DArray, value int) {
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

func peek(arr *DArray) int {
	arrSize := len(arr.data)
	if arrSize == 0 {
		panic("Array is empty")
	}

	return arr.data[arrSize-1]
}

func isEmpty(arr *DArray) bool {
	return len(arr.data) == 0
}

func pop(arr *DArray) int {
	arrSize := len(arr.data)
	if arrSize == 0 {
		panic("Array is empty")
	}

	newSize := arrSize - 1
	value := arr.data[newSize]
	arr.data = arr.data[:newSize]
	if cap(arr.data) > 0 && newSize < cap(arr.data)/4 {
		arrCap := max(cap(arr.data)/2, 1)
		newArr := make([]int, newSize, arrCap)
		copy(newArr, arr.data)
		arr.data = newArr
	}
	return value
}

func main() {
	dynamicArray := DArray{
		data: make([]int, 0, 0),
	}

	fmt.Println("is empty:", isEmpty(&dynamicArray))

	push(&dynamicArray, 1)
	fmt.Println(dynamicArray, cap(dynamicArray.data))

	fmt.Println("is empty:", isEmpty(&dynamicArray))

	push(&dynamicArray, 2)
	push(&dynamicArray, 3)
	push(&dynamicArray, 4)

	fmt.Println("peek:", peek(&dynamicArray))

	push(&dynamicArray, 5)
	push(&dynamicArray, 6)
	push(&dynamicArray, 1)
	push(&dynamicArray, 2)
	push(&dynamicArray, 3)

	fmt.Println("peek:", peek(&dynamicArray))

	push(&dynamicArray, 4)
	push(&dynamicArray, 5)
	push(&dynamicArray, 6)

	fmt.Println(dynamicArray, cap(dynamicArray.data))

	value := pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))
	value = pop(&dynamicArray)
	fmt.Println(dynamicArray, cap(dynamicArray.data))

	fmt.Println(value)

}
