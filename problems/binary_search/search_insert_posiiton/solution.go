package search_insert_posiiton

func SearchInsert(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
