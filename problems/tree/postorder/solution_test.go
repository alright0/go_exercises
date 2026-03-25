package postorder

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreePostorder(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := Postorder(tree.Root)
	target := 7
	if len(result) != target {
		t.Errorf("Result Len %d != %d", len(result), target)
	}

	targetList := []int{4, 3, 2, 3, 4, 2, 1}
	for i := 0; i < len(result); i++ {
		if result[i] != targetList[i] {
			t.Errorf("Result %d != %d", result, targetList)
			break
		}
	}
}

func TestTreePostorderEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := Postorder(tree.Root)
	target := 0
	if len(result) != target {
		t.Errorf("Result Len %d != %d", len(result), target)
	}
}

func TestTreePostorderSortValues(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)

	// Postorder может вернуть отсортированное дерево по возрастанию,
	// если дерево является бинарным деревом поиска
	result := Postorder(tree.Root)
	target := 4
	if len(result) != target {
		t.Errorf("Result Len %d != %d", len(result), target)
	}

	targetList := []int{1, 2, 3, 4}
	for i := 0; i < len(result); i++ {
		if result[i] != targetList[i] {
			t.Errorf("Result %d != %d", result, targetList)
			break
		}
	}
}
