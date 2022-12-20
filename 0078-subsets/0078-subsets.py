class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        return [[]] + self.subsets_recursive(nums)
        
    def subsets_recursive(self, nums: List[int]) -> List[List[int]]:     
        if len(nums) == 0:
            return []
        
        all_subsets = []
        for i in range(len(nums)):
            with_i = self.subsets_recursive(nums[i+1:])
            for elem in with_i:
                elem.insert(0, nums[i])
                
            with_i.insert(0, [nums[i]])    
            all_subsets.extend(with_i)
        
        return all_subsets
        