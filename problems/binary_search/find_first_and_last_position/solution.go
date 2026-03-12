package find_first_and_last_position

func SearchRange(arr []int, target int) [2]int {
	if len(arr) == 0 {
		return [2]int{-1, -1}
	}

	first := findFirst(arr, target)
	last := findLast(arr, target)
	return [2]int{first, last}
}

func findFirst(arr []int, target int) int {
	result := -1
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}
	return result
}

func findLast(arr []int, target int) int {
	result := -1
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}
	return result
}
