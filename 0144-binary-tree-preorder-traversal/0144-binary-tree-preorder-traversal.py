# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
"""
class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if root == None:
            return []
        return [root.val] + self.preorderTraversal(root.left) + self.preorderTraversal(root.right)
"""

class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        stack = [root]
        preorder = []
        while stack:
            n = stack.pop()
            if n is None:
                continue
            preorder.append(n.val)
            stack.append(n.right)
            stack.append(n.left)    
        return preorder    