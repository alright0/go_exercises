package is_symmetric

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeIsSymmetric(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := IsSymmetric(tree.Root)
	if !result {
		t.Errorf("IsSymmetric FAILED. Tree is symmetric")
	}
}

func TestTreeIsSymmetricEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := IsSymmetric(tree.Root)
	if !result {
		t.Errorf("IsSymmetric FAILED. Empty tree is symmetric")
	}
}

func TestTreeNotIsSymmetricTree(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	result := IsSymmetric(tree.Root)
	if result {
		t.Errorf("IsSymmetric FAILED. Tree is not symmetric")
	}
}
