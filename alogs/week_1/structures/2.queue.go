package main

import "fmt"

type QDynamicArray struct {
	data []int
}

func (d *QDynamicArray) Push(value int) {
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

func (d *QDynamicArray) Peek() int {
	arrSize := len(d.data)
	return d.data[arrSize-1]
}

func (d *QDynamicArray) Pop() int {
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

type QStack struct {
	arr QDynamicArray
}

func (s *QStack) Push(value int) {
	s.arr.Push(value)
}

func (s *QStack) Peek() int {
	if s.IsEmpty() {
		panic("Array is empty")
	}
	return s.arr.Peek()
}

func (s *QStack) IsEmpty() bool {
	return len(s.arr.data) == 0
}

func (s *QStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.arr.Pop(), true
}

type Queue struct {
	inStack  QStack
	outStack QStack
}

func (q *Queue) Enqueue(value int) {
	q.inStack.Push(value)
}

func (q *Queue) Dequeue() (int, bool) {
	if val, ok := q.outStack.Pop(); ok {
		return val, true
	}

	for val, ok := q.inStack.Pop(); ok; val, ok = q.inStack.Pop() {
		q.outStack.Push(val)
	}

	return q.outStack.Pop()
}

func main() {

	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)

	//isE := stack.IsEmpty()
	//fmt.Println("is empty:", isE)
	//
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(4)
	//stack.Push(9)
	//
	//fmt.Println(stack)
	//
	//stack.pop()
	//fmt.Println("after pop:", stack)
	//
	//val := stack.Peek()
	//fmt.Println("peek:", val)
	//
	//isE = stack.IsEmpty()
	//fmt.Println("is empty:", isE)

}
