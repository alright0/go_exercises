package validate_bst

import (
	ds "learning/internal/ds/tree"
	"math"
)

func ValidateBstDiapasone(root *ds.TreeNode) bool {
	maxVal := math.MaxInt64
	minVal := math.MinInt64

	return validate(root, minVal, maxVal)
}

func validate(node *ds.TreeNode, minVal int, maxVal int) bool {
	if node == nil {
		return true
	}
	if node.Value < minVal || node.Value > maxVal {
		return false
	}

	return validate(node.Left, minVal, node.Value) && validate(node.Right, node.Value, maxVal)
}
