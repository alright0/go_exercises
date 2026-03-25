package count_complete_tree_nodes

import ds "learning/internal/ds/tree"

// CountNodes https://leetcode.com/problems/count-complete-tree-nodes/
func CountNodes(root *ds.TreeNode) int {
	return _count(root)
}

func _count(node *ds.TreeNode) int {
	if node == nil {
		return 0
	}

	left := _count(node.Left)
	right := _count(node.Right)

	return 1 + left + right
}
