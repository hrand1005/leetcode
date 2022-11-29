# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        if len(nums) == 0:
            return None
        
        median_index = len(nums)//2
        node = TreeNode(
            val=nums[median_index]
        )
        
        node.left = self.sortedArrayToBST(nums[:median_index])
        node.right = self.sortedArrayToBST(nums[median_index+1:])
        
        return node
            