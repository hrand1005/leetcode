# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        smallest = []
        self.kth_smallest_recursive(root, k, smallest)
        return smallest[-1]
    
    def kth_smallest_recursive(self, root: Optional[TreeNode], k: int, smallest: List[int]):
        if root == None:
            return 
        if len(smallest) < k:
            self.kth_smallest_recursive(root.left, k, smallest)
        if len(smallest) < k:
            smallest.append(root.val)
        if len(smallest) < k:
            self.kth_smallest_recursive(root.right, k, smallest)
        return 
            
