# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        max_left = self.max_depth(root.left)
        max_right = self.max_depth(root.right)
        
        diam_left = self.diameterOfBinaryTree(root.left)
        diam_right = self.diameterOfBinaryTree(root.right)
        
        return max(max_left + max_right, diam_left, diam_right)
    
    def max_depth(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        
        left_depth = self.max_depth(root.left)
        right_depth = self.max_depth(root.right)
        
        return max(left_depth, right_depth) + 1
        
        