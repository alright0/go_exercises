package bst_find_max

import ds "learning/internal/ds/tree"

func FindMaxRecursion(node *ds.TreeNode) *ds.TreeNode {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return FindMaxRecursion(node.Right)
	}
	return node
}
