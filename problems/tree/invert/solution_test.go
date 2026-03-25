package invert

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeInvert(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	resultNode := Invert(tree.Root)

	if resultNode.Left.Value != 3 && resultNode.Right.Value != 2 {
		t.Errorf("Tree not inverted: ")
	}
}

func TestTreeInvertEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	resultNode := Invert(tree.Root)
	if resultNode != nil {
		t.Errorf("Result node not empty!")
	}
}
