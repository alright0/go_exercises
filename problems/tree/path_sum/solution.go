package path_sum

import ds "learning/internal/ds/tree"

// PathSum https://leetcode.com/problems/path-sum/
func PathSum(root *ds.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, targetSum)
}

func pathSum(node *ds.TreeNode, targetSum int) bool {
	if node == nil {
		return false
	}

	targetSum -= node.Value
	if targetSum == 0 && node.Left == nil && node.Right == nil {
		return true
	}

	left := pathSum(node.Left, targetSum)
	right := pathSum(node.Right, targetSum)

	return left || right
}
