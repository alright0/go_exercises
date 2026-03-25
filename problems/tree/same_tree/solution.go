package same_tree

import (
	week_5 "learning/internal/ds/tree"
)

func IsSame(lRoot, rRoot *week_5.TreeNode) bool {
	return _isSame(lRoot, rRoot)
}

func _isSame(p, q *week_5.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Value != q.Value {
		return false
	}

	left := _isSame(p.Left, q.Left)
	right := _isSame(p.Right, q.Right)

	return left && right
}
