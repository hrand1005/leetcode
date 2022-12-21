# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.isValidBSTRecursive(root, None, None)
    
    def isValidBSTRecursive(self, root: Optional[TreeNode], min_val, max_val) -> bool:
        if root == None:
            return True
        
        if min_val != None:
            if root.val <= min_val:
                return False
        if max_val != None:
            if max_val <= root.val:
                return False
        
        return self.isValidBSTRecursive(root.left, min_val, root.val) and self.isValidBSTRecursive(root.right, root.val, max_val)
        
        