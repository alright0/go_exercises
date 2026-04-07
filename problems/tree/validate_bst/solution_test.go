package validate_bst

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestValidateBstRecursion(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := ValidateBstRecursion(tree.Root)
	if !result {
		t.Errorf("ValidateBstRecursion FAILED! Tree is valid")
		return
	}
}

func TestValidateBstRecursionEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := ValidateBstRecursion(tree.Root)
	if !result {
		t.Errorf("ValidateBstRecursion FAILED for empty tree!")
	}
}

func TestValidateBstRecursionNotValidTree(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	NodeToPatch := tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	NodeToPatch.Value = 1
	result := ValidateBstRecursion(tree.Root)
	if result {
		t.Errorf("ValidateBstRecursion FAILED! Tree is not valid")
		return
	}
}

func TestValidateBstDiapasone(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := ValidateBstDiapasone(tree.Root)
	if !result {
		t.Errorf("ValidateBstDiapasone FAILED! Tree is valid")
		return
	}
}

func TestValidateBstDiapasoneEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := ValidateBstDiapasone(tree.Root)
	if !result {
		t.Errorf("ValidateBstDiapasone FAILED for empty tree!")
	}
}

func TestValidateBstDiapasoneNotValidTree(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	NodeToPatch := tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	NodeToPatch.Value = 1
	result := ValidateBstDiapasone(tree.Root)
	if result {
		t.Errorf("ValidateBstDiapasone FAILED! Tree is not valid")
		return
	}
}
