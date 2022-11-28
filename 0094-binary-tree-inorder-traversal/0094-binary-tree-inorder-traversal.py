# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
"""
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if root == None:
            return []
        
        return self.inorderTraversal(root.left) + [root.val] + self.inorderTraversal(root.right)
    
"""    
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        inorder = []
        stack = [(root, False)]
        while stack:
            node, visited = stack.pop()
            if node:
                if visited:
                    inorder.append(node.val)
                else:    
                    stack.append((node.right, False))
                    stack.append((node, True))
                    stack.append((node.left, False))

        return inorder