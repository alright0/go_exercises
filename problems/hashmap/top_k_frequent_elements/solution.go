package top_k_frequent_elements

func TopKFrequentElements(arr []int, k int) []int {
	numbers := make(map[int][]int, len(arr))
	buckets := make([][]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if _, ok := numbers[val]; !ok {
			numbers[val] = []int{}
		}
		numbers[val] = append(numbers[val], i)
	}

	for num, freq := range numbers {
		index := len(freq)
		buckets[index] = append(buckets[index], num)
	}

	var topNums []int
	for i := len(buckets) - 1; i >= 0 && len(topNums) < k; i-- {
		for _, num := range buckets[i] {
			topNums = append(topNums, num)
			if len(topNums) == k {
				break
			}
		}
	}

	return topNums
}
