package search_in_rotated_sorted_array

func SearchInRotatedSortedArray(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[left] <= arr[mid] {
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
