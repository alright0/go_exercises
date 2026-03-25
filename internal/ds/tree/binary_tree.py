from typing import Optional


class TreeNode:
    def __init__(
        self, value: int,
        left: Optional["TreeNode"] = None,
        right: Optional["TreeNode"] = None
    ):
        self.value = value
        self.left = left
        self.right = right

class BinaryTree:

    def __init__(self):
        self.root: Optional[TreeNode] = None

    def is_empty(self):
        return not self.root

    def insert(self, value: int) -> TreeNode:
        node = TreeNode(value)
        if self.is_empty():
            self.root = node
            return node

        queue = [self.root]

        while len(queue) > 0:
            current_node = queue.pop(0)

            if not current_node.left:
                current_node.left = node
                break
            queue.append(current_node.left)
            if not current_node.right:
                current_node.right = node
                break
            queue.append(current_node.right)
        return node

    def render(self) -> None:
        """Tree render for BFS only
              1
          2       3
        3   3 ...
        """
        if not self.root:
            print("[]")
            return

        levels = self.level_order_traversal
        height = len(levels)
        for i, level in enumerate(levels):
            gap = 2 ** (height - i - 1)
            line = ""
            for j, val in enumerate(level):
                if j == 0:
                    line += " " * (gap - 1)
                line += str(val)
                line += " " * (gap * 2 - 1)
            print(line.rstrip())

    @property
    def level_order_traversal(self) -> list[list[int]]:
        if not self.root:
            return []

        queue = [self.root]
        result = []
        while len(queue) > 0:
            level_size = len(queue)
            level = []

            for l in range(level_size):
                node = queue.pop(0)
                if node:
                    level.append(node.value)
                    queue.append(node.left)
                    queue.append(node.right)
            result.append(level)
        return result


if __name__ == '__main__':
    tree = BinaryTree()
    tree.insert(1)
    tree.insert(2)
    tree.insert(3)
    tree.insert(3)
    tree.insert(3)
    tree.insert(2)
    tree.insert(2)

    tree.render()