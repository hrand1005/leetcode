"""
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        if k == len(nums):
            return min(nums)
        
        left, equal, right = [], [], []
        pivot = nums[0]
        for n in nums:
            if n == pivot:
                equal.append(n)
            elif n < pivot:
                left.append(n)
            else:
                right.append(n)
        
        if len(right) < k <= len(right) + len(equal):
            return equal[0]
        
        if k <= len(right):
            return self.findKthLargest(right, k)
        else:
            return self.findKthLargest(left, k - len(right) - len(equal))
"""

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        return self.quick_select(nums, k, 0, len(nums)-1)
    
    def quick_select(self, nums: List[int], k: int, left: int, right: int) -> int:
        if k == len(nums[left:right+1]):
            return min(nums[left:right+1])
        
        partition = self.partition(nums, left, right)
        
        if k == len(nums[partition:right+1]):
            return nums[partition]
        elif k <= len(nums[partition+1:right+1]):
            return self.quick_select(nums, k, partition+1, right)
        return self.quick_select(nums, k-len(nums[partition:right+1]), left, partition-1)
   
    def partition(self, nums: List[int], left: int, right: int) -> int:
        l, r = left, right
        pivot = nums[left]
        l += 1
        
        while l <= r:
            while l < right+1 and nums[l] <= pivot:
                l += 1
            while pivot < nums[r]:
                r -= 1
            if l < r:
                nums[l], nums[r] = nums[r], nums[l]
        
        nums[left], nums[l-1] = nums[l-1], nums[left]
        return l-1