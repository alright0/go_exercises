package same_tree

import (
	ds "learning/internal/ds/tree"
	"testing"
)

func TestTreesIsSame(t *testing.T) {
	lTree := ds.BinaryTree{}
	lTree.Insert(1)
	lTree.Insert(2)
	lTree.Insert(2)

	rTree := ds.BinaryTree{}
	rTree.Insert(1)
	rTree.Insert(2)
	rTree.Insert(2)

	result := IsSame(lTree.Root, rTree.Root)
	if !result {
		t.Errorf("IsSame FAILED. Trees is the same")
	}
}

func TestTreesEmptyTreeIsSame(t *testing.T) {
	lTree := ds.BinaryTree{}
	rTree := ds.BinaryTree{}

	result := IsSame(lTree.Root, rTree.Root)
	if !result {
		t.Errorf("IsSame FAILED. Trees is the same")
	}
}

func TestTreesNotIsSame(t *testing.T) {
	lTree := ds.BinaryTree{}
	lTree.Insert(1)
	lTree.Insert(2)
	lTree.Insert(2)

	rTree := ds.BinaryTree{}
	rTree.Insert(1)
	rTree.Insert(2)
	rTree.Insert(3)

	result := IsSame(lTree.Root, rTree.Root)
	if result {
		t.Errorf("IsSame FAILED. Trees is not the same")
	}
}
