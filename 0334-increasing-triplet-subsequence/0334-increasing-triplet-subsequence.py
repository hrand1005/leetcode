"""
class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        if len(nums) < 3:
            return False
        
        maxes = set()
        mins = set()
        
        for i in range(len(nums)):
            for j in range(len(nums)-1, -1, -1):
                if j-1==i:
                    break
                if nums[i] < nums[j]-1:
                    for k in range(i+1, j):
                        if nums[i] < nums[k] < nums[j]:
                            return True
        return False                
        
class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        if len(nums) < 3:
            return False
        start, end = 0, len(nums)-1
        while nums[start+1] <= nums[start]:
            start += 1
            if start == end-1:
                return False
        while nums[end] <= nums[end-1]:
            end -= 1
            if end == start+1:
                return False
        
        for i in range(start+1, end):
            left, right = start, end
            while left < i and nums[i] <= nums[left]:
                left += 1
            while i < right and nums[right] <= nums[i]:
                right -= 1
            if left < i < right:
                return True
        return False        
"""

class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        smallest, second_smallest = float('inf'), float('inf')
        for n in nums:
            if n <= smallest:
                smallest = n
            elif n <= second_smallest:    
                second_smallest = n
            else:
                return True
        return False    
        