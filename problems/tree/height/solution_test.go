package height

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeHeight(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := Height(tree.Root)
	target := 3
	if result != target {
		t.Errorf("Height %d != %d", result, target)
	}
}

func TestTreeHeightEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := Height(tree.Root)
	target := 0
	if result != target {
		t.Errorf("Height %d != %d", result, target)
	}
}
