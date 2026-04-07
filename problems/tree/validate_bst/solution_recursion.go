package validate_bst

import (
	ds "learning/internal/ds/tree"
)

func ValidateBstRecursion(root *ds.TreeNode) bool {
	var prev *int
	return inorder(root, &prev)
}

func inorder(node *ds.TreeNode, prev **int) bool {
	if node == nil {
		return true
	}
	if !inorder(node.Left, prev) {
		return false
	}
	if *prev != nil && node.Value <= **prev {
		return false
	}
	*prev = &node.Value

	return inorder(node.Right, prev)
}

//func ValidateBst(root *ds.TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	minVal := math.MinInt
//	maxVal := math.MaxInt
//
//	node := root
//	for minVal < maxVal {
//		if node.Left != nil {
//			node = node.Left
//			continue
//		}
//		if minVal < node.Value {
//			minVal = node.Value
//		} else {
//			return false
//		}
//		if node.Right != nil {
//			node = node.Right
//			continue
//		}
//	}
//	return true
//}
