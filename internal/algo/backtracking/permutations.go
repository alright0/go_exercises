package backtracking

// Permutations https://leetcode.com/problems/permutations/description/
func Permutations(arr []int) [][]int {
	var result [][]int

	used := make([]bool, len(arr), len(arr))

	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == len(arr) {
			tmp := make([]int, len(path), len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(arr); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, arr[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]int{})
	return result
}
