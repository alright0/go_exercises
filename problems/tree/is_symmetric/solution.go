package is_symmetric

import ds "learning/internal/ds/tree"

func IsSymmetric(root *ds.TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left *ds.TreeNode, right *ds.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Value != right.Value {
		return false
	}

	isMirrorLeft := isMirror(left.Left, right.Right)
	isMirrorRight := isMirror(left.Right, right.Left)

	return isMirrorLeft && isMirrorRight
}
