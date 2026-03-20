package backtracking

// TotalQueens https://leetcode.com/problems/n-queens-ii/description/
func TotalQueens(n int) int {
	var result int
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}
	var path []int

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			result++
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}
			path = append(path, col)
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			backtrack(row + 1)

			diag2[row+col] = false
			diag1[row-col] = false
			cols[col] = false
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
