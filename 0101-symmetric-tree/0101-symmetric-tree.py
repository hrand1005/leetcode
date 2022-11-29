def checkMirrored(node1: Optional[TreeNode], node2: Optional[TreeNode]) -> bool:
    if node1 == None and node2 == None:
        return True
    
    if node1 == None or node2 == None:
        return False
    
    if node1.val != node2.val:
        return False
    
    return checkMirrored(node1.left, node2.right) and checkMirrored(node1.right, node2.left)

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        return checkMirrored(root.left, root.right)