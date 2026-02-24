package ds

import (
	ds "learning/internal/ds/stack"
)

type Queue struct {
	inStack  ds.Stack
	outStack ds.Stack
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
