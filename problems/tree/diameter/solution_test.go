package diameter

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeDiameter(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := Diameter(tree.Root)
	target := 4
	if result != target {
		t.Errorf("Diameter %d != %d", result, target)
	}
}

func TestTreeDiameterEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := Diameter(tree.Root)
	target := 0
	if result != target {
		t.Errorf("diameter %d != %d", result, target)
	}
}
