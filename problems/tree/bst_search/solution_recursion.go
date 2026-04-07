package bst_search

import ds "learning/internal/ds/tree"

func BstSearchRecursion(root *ds.TreeNode, value int) *ds.TreeNode {
	return searchRecursion(root, value)
}

func searchRecursion(node *ds.TreeNode, value int) *ds.TreeNode {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}
	if node.Value > value {
		return searchRecursion(node.Left, value)
	}
	return searchRecursion(node.Right, value)
}
