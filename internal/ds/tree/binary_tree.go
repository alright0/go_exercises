package ds

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (b *BinaryTree) IsEmpty() bool {
	return b.Root == nil
}

// Insert BFS - вставка через поиск в ширину
func (b *BinaryTree) Insert(value int) *TreeNode {
	node := &TreeNode{Value: value}
	if b.IsEmpty() {
		b.Root = node
		return node
	}
	var queue []*TreeNode

	queue = append(queue, b.Root)

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if currentNode.Left == nil {
			currentNode.Left = node
			break
		}
		queue = append(queue, currentNode.Left)

		if currentNode.Right == nil {
			currentNode.Right = node
			break
		}
		queue = append(queue, currentNode.Right)
	}
	return node
}
