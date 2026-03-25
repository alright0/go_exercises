package lowest_common_ancestor

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeLowestCommonAncestor(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	p := tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	q := tree.Insert(4)

	resultNode := LowestCommonAncestor(p, q, tree.Root)
	targetValue := 1
	if resultNode.Value != targetValue {
		t.Errorf("LowestCommonAncestor FAILED. %d != %d", resultNode.Value, targetValue)
	}
}
