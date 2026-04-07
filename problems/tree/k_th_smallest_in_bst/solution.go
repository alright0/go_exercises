package k_th_smallest_in_bst

import ds "learning/internal/ds/tree"

func KthSmallestInBst(root *ds.TreeNode, k int) int {
	result, _ := inorder(root, &k)
	return result
}

func inorder(node *ds.TreeNode, k *int) (int, bool) {
	if node == nil {
		return 0, false
	}

	left, found := inorder(node.Left, k)
	if found {
		return left, true
	}
	(*k)--
	if *k == 0 {
		return node.Value, true
	}
	return inorder(node.Right, k)
}
