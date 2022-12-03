"""
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        
        majority = (len(nums) // 2) + 1
            
        occurrences = {}
        
        for i in range(len(nums)):
            if occurrences.get(nums[i]) == None:
                occurrences[nums[i]] = 1
            else:
                occurrences[nums[i]] += 1
            
            if occurrences[nums[i]] == majority:
                return nums[i]

        return None   
"""

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        num, counter = nums[0], 0
        
        for n in nums:
            if n == num:
                counter += 1
            elif counter > 1:
                counter -= 1 
            else:
                num, counter = n, 1
            
        return num      