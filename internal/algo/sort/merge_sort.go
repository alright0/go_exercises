package sort

type User struct {
	Name string
	Age  int
}

func mergeSort(arr []User) []User {
	if len(arr) <= 1 {
		return arr
	}

	halfLen := len(arr) / 2
	arr1 := mergeSort(arr[:halfLen])
	arr2 := mergeSort(arr[halfLen:])

	return merge(arr1, arr2)
}

func merge(left []User, right []User) []User {
	result := make([]User, 0, len(left)+len(right))
	var i int
	var j int

	for i < len(left) && j < len(right) {
		if left[i].Age <= right[j].Age {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
