package ds

import (
	ds "learning/internal/ds/array"
)

type Stack struct {
	arr ds.DynamicArray
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

func (s *Stack) Length() int {
	return s.arr.Length()
}

func (s *Stack) IsEmpty() bool {
	return s.arr.Length() == 0
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.arr.Pop(), true
}
