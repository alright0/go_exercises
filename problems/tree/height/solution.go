package height

import ds "learning/internal/ds/tree"

func Height(root *ds.TreeNode) int {
	result := _height(root)
	return result
}

func _height(node *ds.TreeNode) int {
	if node == nil {
		return 0
	}

	left := _height(node.Left)
	right := _height(node.Right)

	return 1 + max(left, right)
}
