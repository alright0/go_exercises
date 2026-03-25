package level_order_traversal

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeLevelOrderTraversal(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := LevelOrderTraversal(tree.Root)
	target := [][]int{{1}, {2, 2}, {4, 3, 3, 4}}
	if len(result) != len(target) {
		t.Errorf("LevelOrderTraversal FAILED. %d != %d", len(result), len(target))

	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != target[i][j] {
				t.Errorf("LevelOrderTraversal FAILED. %d != %d", result, target)
			}
		}
	}
}

func TestTreeIsSymmetricEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := LevelOrderTraversal(tree.Root)
	if len(result) != 0 {
		t.Errorf("LevelOrderTraversal FAILED. Empty tree has 0 len")
	}
}
