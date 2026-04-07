package top_k_frequent_elements

import "errors"

type Item struct {
	Value     int
	Frequency int
}

type minHeap []Item

func (h *minHeap) Len() int {
	return len(*h)
}

func (h minHeap) Less(i int, j int) bool {
	return h[i].Frequency < h[j].Frequency
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(item *Item) {
	*h = append(*h, *item)
	h.heapifyUp(len(*h) - 1)
}

func (h *minHeap) Pop() (*Item, error) {
	if (*h).IsEmpty() {
		return nil, errors.New("heap is empty")
	}
	minElement := (*h)[0]
	lastIndex := len(*h) - 1

	(*h)[0] = (*h)[lastIndex]
	*h = (*h)[:lastIndex]

	if !(*h).IsEmpty() {
		h.heapifyDown(0)
	}
	return &minElement, nil
}

func (h *minHeap) heapifyDown(index int) {
	size := len(*h)
	for index < size-1 {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < size && (*h)[left].Frequency < (*h)[smallest].Frequency {
			smallest = left
		}
		if right < size && (*h)[right].Frequency < (*h)[smallest].Frequency {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.Swap(index, smallest)
		index = smallest
	}
}

func (h *minHeap) IsEmpty() bool {
	return len(*h) == 0
}

func (h minHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h[parentIndex].Frequency <= h[index].Frequency {
			break
		}
		h.Swap(index, parentIndex)
		index = parentIndex
	}
}

func TopKFrequentElements(nums []int, k int) []int {
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	h := &minHeap{}

	for value, frequency := range freq {
		item := Item{Value: value, Frequency: frequency}
		h.Push(&item)
		if h.Len() > k {
			h.Pop()
		}
	}
	result := make([]int, 0, k)
	for !h.IsEmpty() {
		item, _ := h.Pop()
		result = append(result, item.Value)
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}
