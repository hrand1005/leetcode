# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
"""
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.isValidBSTRecursive(root, None, None)
    
    def isValidBSTRecursive(self, root: Optional[TreeNode], min_val, max_val) -> bool:
        if root == None:
            return True
        
        if min_val != None and root.val <= min_val:
                return False
        if max_val != None and max_val <= root.val:
                return False
        
        left_valid =  self.isValidBSTRecursive(root.left, min_val, root.val)
        right_valid = self.isValidBSTRecursive(root.right, root.val, max_val)
        
        return left_valid and right_valid
"""
# Inorder traversal
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        inorder_values = []
        self.visit_inorder(root, inorder_values)
        
        for i in range(1, len(inorder_values)):
            if inorder_values[i] <= inorder_values[i-1]:
                return False
        return True
    
    def visit_inorder(self, root: Optional[TreeNode], visited: List[int]):
        if root == None:
            return
        
        self.visit_inorder(root.left, visited)
        visited.append(root.val)
        self.visit_inorder(root.right, visited)
        
        return