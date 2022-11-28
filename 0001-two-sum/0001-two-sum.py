class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        num_to_idx = {}
        for i in range(0, len(nums)):
            comp_idx = num_to_idx.get(target - nums[i])
            if comp_idx != None:
                return [comp_idx, i]
            
            num_to_idx[nums[i]] = i
