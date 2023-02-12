class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        val_to_idx = {}
        for i in range(len(nums)):
            comp = target - nums[i]
            if val_to_idx.get(comp) != None:
                return [val_to_idx[comp], i]
            val_to_idx[nums[i]] = i
        return []    
