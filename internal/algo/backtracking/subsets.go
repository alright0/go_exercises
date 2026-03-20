package backtracking

// Subsets https://leetcode.com/problems/subsets/description/
func Subsets(nums []int) [][]int {
	var result [][]int

	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, []int{})
	return result
}
