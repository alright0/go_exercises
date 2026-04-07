package bst_search

import ds "learning/internal/ds/tree"

func BstSearchCycle(root *ds.TreeNode, value int) *ds.TreeNode {
	node := root

	for node != nil {
		if node.Value == value {
			return node
		}

		if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
