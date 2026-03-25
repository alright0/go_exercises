from typing import Optional

from internal.ds.tree.binary_tree import TreeNode, BinaryTree

# https://leetcode.com/problems/path-sum/
def has_path_sum(root: Optional[TreeNode], target_sum: int) -> bool:
    return _has_path_sum(root, target_sum)

def _has_path_sum(node, target_sum) -> bool:
    if not node:
        return False

    target_sum -= node.value
    if target_sum == 0 and not node.left and not node.right:
        return True

    left = _has_path_sum(node.left, target_sum)
    right = _has_path_sum(node.right, target_sum)

    return left or right



if __name__ == '__main__':
    tree = BinaryTree()
    tree.insert(1)
    tree.insert(2)
    tree.insert(2)
    tree.insert(4)
    tree.insert(3)
    tree.insert(3)
    tree.insert(4)

    assert has_path_sum(tree.root, 7)
    assert has_path_sum(tree.root, 6)
    assert not has_path_sum(tree.root, 2)

    tree = BinaryTree()
    tree.insert(1)
    tree.insert(2)

    assert not has_path_sum(tree.root, 1)
