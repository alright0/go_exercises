package path_sum

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreePathSum(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := PathSum(tree.Root, 7)
	if !result {
		t.Errorf("LowestCommonAncestor FAILED. %t != %t", result, true)
	}
}
