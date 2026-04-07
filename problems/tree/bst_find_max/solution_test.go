package bst_find_max

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestBstFindMax(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := FindMaxRecursion(tree.Root)
	target := 17
	if result == nil {
		t.Errorf("FindMax FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMax FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestBstFindMaxRecursionDegenerate(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)

	result := FindMaxRecursion(tree.Root)
	target := 4
	if result == nil {
		t.Errorf("FindMax FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMax FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestFindMaxRecursionEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := FindMaxRecursion(tree.Root)
	target := 0
	if result != nil {
		t.Errorf("FindMax FAILED! %d != %d", result.Value, target)
	}
}

func TestBstFindMaxIterate(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := FindMaxIterate(tree.Root)
	target := 17
	if result == nil {
		t.Errorf("FindMax FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMax FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestBstFindMaxIterateDegenerate(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)

	result := FindMaxIterate(tree.Root)
	target := 4
	if result == nil {
		t.Errorf("FindMax FAILED! result should not be nil")
		return
	}
	if result.Value != target {
		t.Errorf("FindMax FAILED! Target min is %d, found %d ", target, result.Value)
	}
}

func TestFindMaxIterateEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := FindMaxIterate(tree.Root)
	target := 0
	if result != nil {
		t.Errorf("FindMax FAILED! %d != %d", result.Value, target)
	}
}
