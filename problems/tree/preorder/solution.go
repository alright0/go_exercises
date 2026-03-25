package preorder

import ds "learning/internal/ds/tree"

func Preorder(root *ds.TreeNode) []int {
	var result []int

	_preorder(root, &result)

	return result
}

func _preorder(node *ds.TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Value)

	_preorder(node.Left, result)
	_preorder(node.Right, result)
}
