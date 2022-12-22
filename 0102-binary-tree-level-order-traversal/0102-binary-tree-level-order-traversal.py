# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root == None:
            return []
        
        values = []
        queue = [(root, 0)]
        
        while queue:
            node, level = queue.pop(0)
            if node == None:
                continue
                
            if len(values) < level + 1:
                values.append([node.val])
            else:
                values[level].append(node.val)
                
            queue.append((node.left, level+1))
            queue.append((node.right, level+1))
        
        return values
