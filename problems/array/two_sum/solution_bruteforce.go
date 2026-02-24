package two_sum

func TwoSumBruteforce(nums []int, target int) (int, int) {
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := indexMap[complement]; ok {
			return index, i
		}
		indexMap[num] = i
	}
	return 0, 0
}
