package backtracking

// NQueens https://leetcode.com/problems/n-queens/description/?
func NQueens(n int) [][]string {
	var result [][]string
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}
	var path []int

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			board := make([]string, n)

			for r, c := range path {
				row := make([]byte, n)
				for i := range row {
					row[i] = '.'
				}
				row[c] = 'Q'
				board[r] = string(row)
			}

			result = append(result, board)
			return
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
