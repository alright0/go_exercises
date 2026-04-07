package ds

import "testing"

func TestBinarySearchTreeInsertOrdered(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(2)

	minimum := 2
	if tree.Root.Left.Left.Value != minimum {
		t.Errorf("Left leaf should be %d but is: %d", minimum, tree.Root.Left.Value)
	}

	maximum := 7
	if tree.Root.Right.Value != maximum {
		t.Errorf("Left value should be %d but is: %d", maximum, tree.Root.Left.Value)

	}
}

func TestBinarySearchTreeInsertNotOrdered(t *testing.T) {
	tree := BinarySearchTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	minimum := 1
	if tree.Root.Value != minimum {
		t.Errorf("Left leaf should be %d but is: %d", minimum, tree.Root.Left.Value)
	}

	maximum := 4
	if tree.Root.Right.Right.Right.Value != maximum {
		t.Errorf("Left value should be %d but is: %d", maximum, tree.Root.Left.Value)

	}
}

func TestBinarySearchTreeDeleteNodeWithoutChildren(t *testing.T) {
	/* Удаление листа(ноды без потомков) 4
	   1                     1
	    \                     \
	     2                     2
	      \       ------>       \
	       3                     3
	        \
	         4
	*/
	tree := BinarySearchTree{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	toDelete := 4
	tree.Delete(toDelete)

	if tree.Root.Right.Right.Right != nil {
		t.Errorf("Node with Value %d not deleted ", toDelete)
	}
}

func TestBinarySearchTreeDeleteNodeWithOneLeftChild(t *testing.T) {
	/* Удаление ноды с одним правым листом 2 -> 1
	       4                   4
	      / \                 / \
	     2   5    ------>    1   5
	    /
	   1
	*/
	tree := BinarySearchTree{}
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)

	toDelete := 2
	candidate := 1
	tree.Delete(toDelete)

	if tree.Root.Left.Value == toDelete {
		t.Errorf("Node with Value %d not deleted ", toDelete)
	}

	if tree.Root.Left.Value != candidate {
		t.Errorf("Node new Value is not %d", candidate)
	}
}

func TestBinarySearchTreeDeleteNodeWithOneRightChild(t *testing.T) {
	/* Удаление ноды с одним правым листом 2 -> 3
	     4                   4
	    / \                 / \
	   2   5    ------>    3   5
	    \
	     3
	*/

	tree := BinarySearchTree{}
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(3)

	toDelete := 2
	candidate := 3
	tree.Delete(toDelete)

	if tree.Root.Left.Value == toDelete {
		t.Errorf("Node with Value %d not deleted ", toDelete)
	}

	if tree.Root.Left.Value != candidate {
		t.Errorf("Node new Value is not %d", candidate)
	}

}

func TestBinarySearchTreeDeleteNodeWithTwoChildren(t *testing.T) {
	/* Удаление ноды с двумя листьями 2 -> 3
	       4                   4
	      / \                 / \
	     2   5    ------>    1   5
	    / \                   \
	   1   3                   3
	*/
	tree := BinarySearchTree{}
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)

	toDelete := 2
	candidate := 3
	tree.Delete(toDelete)

	if tree.Root.Left.Value == toDelete {
		t.Errorf("Node with Value %d not deleted ", toDelete)
	}
	if tree.Root.Left.Value != candidate {
		t.Errorf("Node new Value is %d not %d", tree.Root.Left.Value, candidate)
	}
}

func TestBinarySearchTreeDeleteNodeWithNestedChildren(t *testing.T) {
	/* Удаление ноды с вложенными листьями 5 -> 6
	       10                     10
	      /  \                   /  \
	     5   15                 6    15
	    /  \       ------>     / \
	   2    7                 2   7
	    \   / \                \   \
	     4 6   8                4   8
	*/
	tree := BinarySearchTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(6)

	toDelete := 5
	candidate := 6
	tree.Delete(toDelete)

	if tree.Root.Left.Value == toDelete {
		t.Errorf("Node with Value %d not deleted ", toDelete)
	}
	if tree.Root.Left.Value != candidate {
		t.Errorf("Node new Value is %d not %d", tree.Root.Left.Value, candidate)
	}
}
