package ds

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func (h *MinHeap) Push(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}
	minElement := h.data[0]
	lastIndex := len(h.data) - 1

	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}
	return minElement, true
}

func (h *MinHeap) Peek() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}
	return h.data[0], true
}

func (h *MinHeap) Size() int {
	return len(h.data)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := parent(index)
		if h.data[parentIndex] <= h.data[index] {
			break
		}
		h.data[index], h.data[parentIndex] = h.data[parentIndex], h.data[index]
		index = parentIndex
	}
}

func (h *MinHeap) heapifyDown(index int) {
	size := len(h.data)
	for index < size-1 {
		left := leftChild(index)
		right := rightChild(index)
		smallest := index

		if left < size && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < size && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}

		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

// Индексы
func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
