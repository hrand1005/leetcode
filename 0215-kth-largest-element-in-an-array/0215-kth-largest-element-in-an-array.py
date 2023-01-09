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