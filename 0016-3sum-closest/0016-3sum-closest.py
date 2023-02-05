"""
class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        closest = nums[0] + nums[1] + nums[2]
        min_diff = abs(target - closest)
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                for k in range(j+1, len(nums)):
                    this = nums[i] + nums[j] + nums[k]
                    if abs(target - this) < min_diff:
                        min_diff = abs(target - this)
                        closest = this
        return closest                
"""

class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        closest = nums[0] + nums[1] + nums[2]
        for i in range(len(nums)-2):
            left, right = i+1, len(nums)-1
            while left < right:
                this = nums[i] + nums[left] + nums[right]
                if this == target:
                    return this
                if abs(target - this) < abs(target - closest):
                    closest = this
                    
                if this < target:
                    left += 1
                else:
                    right -= 1
        return closest            