package bst_height

import ds "learning/internal/ds/tree"

func BstHeight(root *ds.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(BstHeight(root.Left), BstHeight(root.Right))
}
