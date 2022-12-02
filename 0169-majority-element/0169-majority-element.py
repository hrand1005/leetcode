class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        
        majority = len(nums) // 2
        occurrences = {}
        
        for i in range(len(nums)):
            if occurrences.get(nums[i]) == None:
                occurrences[nums[i]] = 1
            elif occurrences[nums[i]] >= majority:
                return nums[i]
            else:
                occurrences[nums[i]] += 1
