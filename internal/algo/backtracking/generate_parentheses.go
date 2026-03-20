package backtracking

// GenerateParentheses https://leetcode.com/problems/generate-parentheses/description/
func GenerateParentheses(n int) []string {
	var result []string

	var backtrack func(path string, open int, close int)
	backtrack = func(path string, open int, close int) {
		if len(path) == 2*n {
			result = append(result, path)
			return
		}

		if open < n {
			backtrack(path+"(", open+1, close)
		}

		if close < open {
			backtrack(path+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}
