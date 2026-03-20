package backtracking

// CombinationSum https://leetcode.com/problems/combination-sum/description
func CombinationSum(arr []int, target int) [][]int {
	var result [][]int

	var backtrack func(start int, path []int, currentSum int)
	backtrack = func(start int, path []int, currentSum int) {
		if currentSum == target {
			result = append(result, append([]int{}, path...))
			return
		}

		if currentSum > target {
			return
		}

		for i := start; i < len(arr); i++ {
			currentSum += arr[i]
			path = append(path, arr[i])
			backtrack(i, path, currentSum)
			path = path[:len(path)-1]
			currentSum -= arr[i]
		}
	}

	backtrack(0, []int{}, 0)
	return result
}
