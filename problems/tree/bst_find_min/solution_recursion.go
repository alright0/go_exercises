package bst_find_min

import ds "learning/internal/ds/tree"

func FindMinRecursion(node *ds.TreeNode) *ds.TreeNode {
	if node == nil {
		return nil
	}

	if node.Left != nil {
		return FindMinRecursion(node.Left)
	}
	return node
}
