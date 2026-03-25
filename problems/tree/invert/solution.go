package invert

import ds "learning/internal/ds/tree"

func Invert(root *ds.TreeNode) *ds.TreeNode {
	return _invert(root)
}

func _invert(node *ds.TreeNode) *ds.TreeNode {
	if node == nil {
		return nil
	}

	node.Left, node.Right = node.Right, node.Left

	_invert(node.Left)
	_invert(node.Right)

	return node
}
