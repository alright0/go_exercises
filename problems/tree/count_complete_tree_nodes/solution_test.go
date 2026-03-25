package count_complete_tree_nodes

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestCountNodes(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)

	result := CountNodes(tree.Root)
	target := 3
	if result != target {
		t.Errorf("CountNodes FAILED! %d != %d", result, target)
	}
}

func TestCountNodesEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := CountNodes(tree.Root)
	target := 0
	if result != target {
		t.Errorf("CountNodes FAILED! %d != %d", result, target)
	}
}
