package level_order_traversal

import ds "learning/internal/ds/tree"

func LevelOrderTraversal(root *ds.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*ds.TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Value)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
