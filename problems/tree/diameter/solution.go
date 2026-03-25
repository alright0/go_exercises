package diameter

import ds "learning/internal/ds/tree"

func Diameter(root *ds.TreeNode) int {
	maxDiameter := 0
	_ = diamDfs(root, &maxDiameter)
	return maxDiameter
}

func diamDfs(node *ds.TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}

	left := diamDfs(node.Left, maxDiameter)
	right := diamDfs(node.Right, maxDiameter)

	if left+right > *maxDiameter {
		*maxDiameter = left + right
	}
	return 1 + max(left, right)
}
