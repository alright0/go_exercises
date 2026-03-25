package is_balanced

import ds "learning/internal/ds/tree"

func IsBalanced(root *ds.TreeNode) bool {
	balance := _isBalanced(root)
	return balance != -1
}

func _isBalanced(node *ds.TreeNode) int {
	if node == nil {
		return 0
	}

	left := _isBalanced(node.Left)
	right := _isBalanced(node.Right)

	if left == -1 || right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
