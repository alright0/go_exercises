package k_th_smallest_in_bst

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestBstKthSmallestInBst(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	k := 3
	result := KthSmallestInBst(tree.Root, k)
	target := 7
	if result != target {
		t.Errorf("KthSmallestInBst FAILED! %d != %d for k=%d", result, target, k)
		return
	}
}

func TestKthSmallestInBstEmptyTree(t *testing.T) {
	tree := ds.BinarySearchTree{}

	result := KthSmallestInBst(tree.Root, 0)
	target := 0
	if result != target {
		t.Errorf("KthSmallestInBst FAILED for empty tree! %d != %d", result, target)
	}
}

func TestBstKthSmallestInBstKIsBiggerThenTreeLen(t *testing.T) {
	tree := ds.BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(17)

	k := 25
	result := KthSmallestInBst(tree.Root, k)
	target := 0
	if result != target {
		t.Errorf("TestBstKthSmallestInBstKIsBiggerThenTreeLen FAILED! %d != %d for k=%d", result, target, k)
		return
	}
}
