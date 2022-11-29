# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
"""
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        return self.checkMirrored(root.left, root.right)
    
    def checkMirrored(self, node1: Optional[TreeNode], node2: Optional[TreeNode]) -> bool:
        if node1 == None and node2 == None:
            return True

        if node1 == None or node2 == None:
            return False

        if node1.val != node2.val:
            return False

        return self.checkMirrored(node1.left, node2.right) and self.checkMirrored(node1.right, node2.left)
"""

class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        if root == None:
            return True
        
        stack = [root.left, root.right]
        
        while stack:
            right = stack.pop()
            left = stack.pop()
            
            if right == None and left == None:
                continue
                
            if right == None or left == None:
                return False
                
            if right.val != left.val:
                return False
            
            stack.extend([left.right, right.left])
            stack.extend([left.left, right.right])
                
        return True
                
                