package lowest_common_ancestor

import ds "learning/internal/ds/tree"

func LowestCommonAncestor(p, q, root *ds.TreeNode) *ds.TreeNode {
	return dfs(root, p, q)
}

func dfs(node, p, q *ds.TreeNode) *ds.TreeNode {
	if node == nil {
		return nil
	}

	if node == p || node == q {
		return node
	}

	left := dfs(node.Left, p, q)
	right := dfs(node.Right, p, q)

	if left != nil && right != nil {
		return node
	}

	if left != nil {
		return left
	}

	return right
}
