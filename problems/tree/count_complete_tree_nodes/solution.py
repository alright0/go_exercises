from typing import Optional

from internal.ds.tree.binary_tree import TreeNode, BinaryTree


# https://leetcode.com/problems/count-complete-tree-nodes/
def count_nodes(root: Optional[TreeNode]) -> int:
    return _count(root)

def _count(node: Optional[TreeNode]) -> int:
    if not node:
        return 0

    return 1 + _count(node.left) + _count(node.right)


if __name__ == '__main__':
    tree = BinaryTree()
    tree.insert(1)
    tree.insert(1)
    tree.insert(1)
    tree.insert(1)

    result = count_nodes(tree.root)
    target = 4
    assert result == target, f"{result} != {target}"

