package is_balanced

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreeIsBalanced(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(3)
	tree.Insert(4)

	result := IsBalanced(tree.Root)
	if !result {
		t.Errorf("IsBalanced FAILED. Tree is balanced")
	}
}

func TestTreeIsBalancedEmptyTree(t *testing.T) {
	tree := ds.BinaryTree{}

	result := IsBalanced(tree.Root)
	if !result {
		t.Errorf("IsBalanced FAILED. Empty tree is balanced")
	}
}

func TestTreeNotIsBalancedTree(t *testing.T) {
	tree := ds.BinaryTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(2)
	node := tree.Insert(4)
	// Вставляем ноду слева, чтобы дерево стало несбалансированным
	node.Left = &ds.TreeNode{Value: 1}

	result := IsBalanced(tree.Root)
	if result {
		t.Errorf("IsBalanced FAILED. Tree is not balanced")
	}
}
