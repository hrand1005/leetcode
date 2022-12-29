# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        queue = [[root]]
        zigzag = []
        level = 0
        
        while queue:
            level_nodes = queue.pop(0)
            new_level, this_zig = [], []
            for n in level_nodes:
                if n == None:
                    continue
                    
                this_zig.append(n.val)
                if level % 2 == 0:    
                    new_level = [n.right, n.left] + new_level
                else:
                    new_level = [n.left, n.right] + new_level
            
            if len(new_level) > 0:
                zigzag.append(this_zig)
                queue.append(new_level)
                level += 1
        
        return zigzag
        
                    
                        
        
        