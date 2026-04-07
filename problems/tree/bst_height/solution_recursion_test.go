package bst_height

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestBstHeight(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := BstHeight(tree.Root)
	target := 17
	if result != target {
		t.Errorf("BstHeight FAILED! %d != %d", result, target)
		return
	}
}

func TestFindMaxEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := BstHeight(tree.Root)
	target := 0
	if result != target {
		t.Errorf("BstHeight FAILED! %d != %d", result, target)
	}
}
