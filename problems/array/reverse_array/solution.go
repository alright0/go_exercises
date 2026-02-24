package reverse_array

func ReverseArray(arr []int) []int {
	newArr := make([]int, len(arr), len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		newArr[len(arr)-i-1] = arr[i]
	}
	return newArr
}
