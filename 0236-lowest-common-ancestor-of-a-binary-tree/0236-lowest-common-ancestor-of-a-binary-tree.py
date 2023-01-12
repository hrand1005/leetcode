# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
"""
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if self.in_subtree(root.left, p) and self.in_subtree(root.left, q):
            return self.lowestCommonAncestor(root.left, p, q)
        elif self.in_subtree(root.right, p) and self.in_subtree(root.right, q):
            return self.lowestCommonAncestor(root.right, p, q)
        return root
    
    def in_subtree(self, root: 'TreeNode', node: 'TreeNode') -> bool:
        if root == None:
            return False
        if root == node:
            return True
        
        in_left = self.in_subtree(root.left, node)
        in_right = self.in_subtree(root.right, node)
        
        return in_left or in_right
"""

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        stack = [(root, [])]
        path_with_p = None
        while stack:
            node, path = stack.pop()
            new_path = path + [node]
            if node == p:
                path_with_p = new_path
                break
            if node != None:
                stack.append((node.left, new_path))
                stack.append((node.right, new_path))
                
        if q in path_with_p:
            return q
        
        while path_with_p:
            candidate = path_with_p.pop()
            if self.in_subtree(candidate, q):
                return candidate
            candidate.left = None
            candidate.right = None
            
    def in_subtree(self, root: 'TreeNode', node: 'TreeNode') -> bool:
        if root == None:
            return False
        if root == node:
            return True
        return self.in_subtree(root.left, node) or self.in_subtree(root.right, node)
