package inorder

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeInorder(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := Inorder(tree.Root)
	target := 7
	if len(result) != target {
		t.Errorf("Result Len %d != %d", len(result), target)
	}

	targetList := []int{4, 2, 3, 1, 3, 2, 4}
	for i := 0; i < len(result); i++ {
		if result[i] != targetList[i] {
			t.Errorf("Result %d != %d", result, targetList)
			break
		}
	}
}

func TestTreeInorderEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := Inorder(tree.Root)
	target := 0
	if len(result) != target {
		t.Errorf("Result Len %d != %d", len(result), target)
	}
}
