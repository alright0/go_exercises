package bst_search

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestBstSearchRecursion(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	target := tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := BstSearchRecursion(tree.Root, 7)
	if result != target {
		t.Errorf("BstSearch FAILED node %d not found ", target.Value)
	}
}

func TestBstSearchRecursionEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := BstSearchRecursion(tree.Root, 1)
	if result != nil {
		t.Errorf("BstSearch in empty tree FAILED. Result must be nil, but got %d", result.Value)
	}
}

func TestBstSearchCycle(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	target := tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	result := BstSearchCycle(tree.Root, 7)
	if result != target {
		t.Errorf("BstSearch FAILED node %d not found ", target.Value)
	}
}

func TestBstSearchCycleEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := BstSearchCycle(tree.Root, 1)
	if result != nil {
		t.Errorf("BstSearch in empty tree FAILED. Result must be nil, but got %d", result.Value)
	}
}
