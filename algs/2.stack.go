package main

import "fmt"

type DynamicArray struct {
	data []int
}

func (d *DynamicArray) Push(value int) {
	arrSize := len(d.data)
	if arrSize == cap(d.data) {
		arrCap := max(cap(d.data)*2, 1)
		newArr := make([]int, arrSize, arrCap)
		copy(newArr, d.data)
		d.data = newArr
	}

	d.data = d.data[:arrSize+1]
	d.data[arrSize] = value
}

func (d *DynamicArray) Peek() int {
	arrSize := len(d.data)
	return d.data[arrSize-1]
}

func (d *DynamicArray) Pop() int {
	arrSize := len(d.data)

	newSize := arrSize - 1
	value := d.data[newSize]
	d.data = d.data[:newSize]
	if cap(d.data) > 0 && newSize < cap(d.data)/4 {
		arrCap := max(cap(d.data)/2, 1)
		newArr := make([]int, newSize, arrCap)
		copy(newArr, d.data)
		d.data = newArr
	}
	return value
}

type Stack struct {
	arr DynamicArray
}

func (s *Stack) Push(value int) {
	s.arr.Push(value)
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Array is empty")
	}
	return s.arr.Peek()
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr.data) == 0
}

func (s *Stack) pop() int {
	if s.IsEmpty() {
		panic("Array is empty")
	}
	return s.arr.Pop()
}

func main() {
	dynamicArray := DynamicArray{
		data: make([]int, 0, 0),
	}

	stack := Stack{
		arr: dynamicArray,
	}

	isE := stack.IsEmpty()
	fmt.Println("is empty:", isE)

	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(9)

	fmt.Println(stack)

	stack.pop()
	fmt.Println("after pop:", stack)

	val := stack.Peek()
	fmt.Println("peek:", val)

	isE = stack.IsEmpty()
	fmt.Println("is empty:", isE)

}
