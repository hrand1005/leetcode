# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        if root == None:
            return True
        if 1 < abs(self.depth(root.left) - self.depth(root.right)):
            return False
        return self.isBalanced(root.left) and self.isBalanced(root.right)
    
    def depth(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        return 1 + max(self.depth(root.left), self.depth(root.right))
