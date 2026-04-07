package bst_find_max

import ds "learning/internal/ds/tree"

func FindMaxIterate(node *ds.TreeNode) *ds.TreeNode {
	for node != nil && node.Right != nil {
		node = node.Right
	}
	return node
}
