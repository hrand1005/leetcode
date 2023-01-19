# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if root is None:
            return False
        stack = [(root, root.val)]
        while stack:
            node, acc = stack.pop()
            if node.left is None and node.right is None:
                if acc == targetSum:
                    return True
            if node.left:
                stack.append((node.left, acc + node.left.val))
            if node.right:
                stack.append((node.right, acc + node.right.val))
        return False        
        