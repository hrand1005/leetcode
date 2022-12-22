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
        
        values = [[root.val]]
        queue = [(root, 0)]
        
        while queue:
            node, level = queue.pop(0)
            if node == None:
                continue
                
            child_vals = []
            if node.left != None:
                child_vals.append(node.left.val)
            if node.right != None:
                child_vals.append(node.right.val)
                
            new_level = level+1
            if len(child_vals) != 0:
                if len(values) < new_level + 1:
                    values.append(child_vals)
                else:    
                    values[new_level].extend(child_vals)    
                
            queue.append((node.left, new_level))
            queue.append((node.right, new_level))
        
        return values
