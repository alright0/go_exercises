package ds

import "testing"

func TestMinHeapPush(t *testing.T) {
	h := NewMinHeap()
	h.Push(5)
	h.Push(3)
	h.Push(8)
	h.Push(1)

	target := []int{1, 3, 8, 5}
	for i := 0; i < len(h.data)-1; i++ {
		if h.data[i] != target[i] {
			t.Errorf("Min Heap Push FAILED, %d != %d", h.data, target)
			return
		}
	}
}

func TestMinHeapPeek(t *testing.T) {
	h := NewMinHeap()
	h.Push(5)
	h.Push(3)
	h.Push(8)
	h.Push(1)

	result, found := h.Peek()
	target := 1
	if result != target || !found {
		t.Errorf("Min Heap Peek FAILED, %d != %d", result, target)
		return
	}
}

func TestMinHeapPeekEmptyHeap(t *testing.T) {
	h := NewMinHeap()

	result, found := h.Peek()
	target := 0
	if result != target || found {
		t.Errorf("Min Heap Peek FAILED, %d != %d", result, target)
		return
	}
}

func TestMinHeapPop(t *testing.T) {
	h := NewMinHeap()
	h.Push(5)
	h.Push(3)
	h.Push(8)
	h.Push(1)

	result, popped := h.Pop()
	target := 1
	if result != target || !popped {
		t.Errorf("Min Heap Pop FAILED, %d != %d", result, target)
		return
	}

	targetArr := []int{3, 5, 8}
	if len(h.data) != len(targetArr) {
		t.Errorf("Min Heap Pop FAILED, %d != %d", h.data, targetArr)
	}

	for i := 0; i < len(h.data); i++ {
		if h.data[i] != targetArr[i] {
			t.Errorf("Min Heap Pop FAILED, %d != %d", h.data, targetArr)
			return
		}
	}
}

func TestMinHeapPopEmpty(t *testing.T) {
	h := NewMinHeap()

	result, popped := h.Pop()
	target := 0
	if result != target || popped {
		t.Errorf("Min Heap Pop on empty heap FAILED, %d != %d", result, target)
		return
	}

	targetArr := []int{}
	if len(h.data) != len(targetArr) {
		t.Errorf("Min Heap Pop FAILED, %d != %d", h.data, targetArr)
	}

}
