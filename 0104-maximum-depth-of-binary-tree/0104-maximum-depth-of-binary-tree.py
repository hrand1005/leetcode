# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
"""
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        
        path_lengths = []
        stack = [(root, 1)]
        while stack:
            node, depth = stack.pop()
            
            if node.left == None and node.right == None:
                path_lengths.append(depth)
                continue
                
            if node.left:
                stack.append((node.left, depth+1))
            
            if node.right:
                stack.append((node.right, depth+1))
        
        return max(path_lengths)
"""   
    
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        return self.recursiveHelper(root, 0)

    def recursiveHelper(self, node: Optional[TreeNode], depth) -> int:
        if node == None:
            return depth
        
        return max(self.recursiveHelper(node.left, depth+1), self.recursiveHelper(node.right, depth+1))