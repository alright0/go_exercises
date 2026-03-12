package sort

func QuickSortL(arr []int, left int, right int) {
	if left < right {
		p := partitionL(arr, left, right)

		QuickSortL(arr, left, p-1)
		QuickSortL(arr, p+1, right)
	}
}

func partitionL(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
