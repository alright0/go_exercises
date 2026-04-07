package bst_find_min

import ds "learning/internal/ds/tree"

func FindMinIterate(node *ds.TreeNode) *ds.TreeNode {
	for node != nil && node.Left != nil {
		node = node.Left
	}
	return node
}
