package k_closest_ponts_to_origin

import (
	"container/heap"
)

type Point struct {
	x, y, distance int
}

func distance(x, y int) int {
	return x*x + y*y
}

type MaxHeap []Point

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(point interface{}) {
	*h = append(*h, point.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TopKClosestPointsToOrigin(arr [][]int, k int) [][]int {
	h := MaxHeap{}
	heap.Init(&h)

	for _, p := range arr {
		point := Point{p[0], p[1], distance(p[0], p[1])}

		heap.Push(&h, point)
		if h.Len() > k {
			heap.Pop(&h)
		}

	}

	result := make([][]int, 0, k)
	for h.Len() > 0 {
		item := heap.Pop(&h).(Point)
		result = append(result, []int{item.x, item.y})
	}

	return result
}
