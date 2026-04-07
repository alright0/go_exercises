package bst_find_min

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestBstFindMinRecursion(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := FindMinRecursion(tree.Root)
	target := 2
	if result == nil {
		t.Errorf("FindMin FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMin FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestBstFindMinRecursionDegenerate(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	result := FindMinRecursion(tree.Root)
	target := 1
	if result == nil {
		t.Errorf("FindMin FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMin FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestFindMinRecursionEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := FindMinRecursion(tree.Root)
	target := 0
	if result != nil {
		t.Errorf("FindMin FAILED! %d != %d", result.Value, target)
	}
}

func TestBstFindMinIterate(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := FindMinIterate(tree.Root)
	target := 2
	if result == nil {
		t.Errorf("FindMin FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMin FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestBstFindMinIterateDegenerate(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	result := FindMinIterate(tree.Root)
	target := 1
	if result == nil {
		t.Errorf("FindMin FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMin FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestFindMinIterateEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := FindMinIterate(tree.Root)
	target := 0
	if result != nil {
		t.Errorf("FindMin FAILED! %d != %d", result.Value, target)
	}
}
