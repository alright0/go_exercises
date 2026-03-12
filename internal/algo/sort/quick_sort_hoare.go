package sort

func QuickSort(arr []int, left int, right int) {
	if left < right {
		p := partitionHoare(arr, left, right)

		QuickSort(arr, left, p)
		QuickSort(arr, p+1, right)
	}
}

func partitionHoare(arr []int, left int, right int) int {
	pivot := arr[(left+right)/2]

	i := left - 1
	j := right + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}
