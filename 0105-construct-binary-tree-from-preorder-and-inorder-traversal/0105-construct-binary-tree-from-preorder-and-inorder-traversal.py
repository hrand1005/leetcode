# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0:
            return None
        
        root = TreeNode(val=preorder.pop(0))
        root_index = inorder.index(root.val)
        
        inorder_left = inorder[:root_index]
        inorder_right = inorder[root_index+1:]
        
        preorder_left = preorder[:len(inorder_left)]
        preorder_right = preorder[len(inorder_left):]
        
        root.left = self.buildTree(preorder_left, inorder_left)
        root.right = self.buildTree(preorder_right, inorder_right)
        
        return root
        
        
