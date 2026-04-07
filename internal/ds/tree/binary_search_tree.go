package ds

type BinarySearchTree struct {
	Root *TreeNode
}

func (b *BinarySearchTree) Delete(value int) {
	if b.Root == nil {
		return
	}
	b.Root = b.delete(b.Root, value)
}

func (b *BinarySearchTree) delete(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.Left = b.delete(root.Left, value)
	} else if value > root.Value {
		root.Right = b.delete(root.Right, value)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil && root.Right != nil {
			root = root.Right
		} else if root.Right == nil && root.Left != nil {
			root = root.Left
		} else {
			successor := b.minimum(root.Right)
			root.Value = successor.Value
			root.Right = b.delete(root.Right, successor.Value)
		}
	}
	return root
}

func (b *BinarySearchTree) minimum(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return b.minimum(root.Left)
	}
	return root
}

func (b *BinarySearchTree) Insert(value int) *TreeNode {
	if b.Root == nil {
		b.Root = &TreeNode{Value: value}
		return b.Root
	}
	return b.insert(b.Root, value)
}

func (b *BinarySearchTree) insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value}
	}

	if value < root.Value {
		root.Left = b.insert(root.Left, value)
	} else {
		root.Right = b.insert(root.Right, value)
	}

	return root
}
