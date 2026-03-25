package postorder

import ds "learning/internal/ds/tree"

func Postorder(root *ds.TreeNode) []int {
	var result []int
	_postorder(root, &result)
	return result
}

func _postorder(node *ds.TreeNode, result *[]int) {
	if node == nil {
		return
	}

	_postorder(node.Left, result)
	_postorder(node.Right, result)
	*result = append(*result, node.Value)
}
