package two_sum

func TwoSumOptimal(arr []int, target int) (int, int) {
	m := map[int]int{}

	for i := 0; i < len(arr); i++ {
		targetNum := target - arr[i]
		if candidateIndex, ok := m[targetNum]; ok {
			return candidateIndex, i
		}
		m[arr[i]] = i
	}
	return 0, 0
}
