package inorder

import (
	week_5_structures "learning/internal/ds/tree"
)

func Inorder(root *week_5_structures.TreeNode) []int {
	var result []int

	_inorder(root, &result)

	return result
}

func _inorder(node *week_5_structures.TreeNode, result *[]int) {
	if node == nil {
		return
	}

	_inorder(node.Left, result)
	*result = append(*result, node.Value)
	_inorder(node.Right, result)
}
