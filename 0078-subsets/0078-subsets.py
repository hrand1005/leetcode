class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        return [[]] + self.subsets_recursive(nums)
        
    def subsets_recursive(self, nums: List[int]) -> List[List[int]]:     
        if len(nums) == 0:
            return []
        
        all_subsets = []
        for i in range(len(nums)):
            sub = self.subsets_recursive(nums[i+1:])
            for elem in sub:
                elem.append(nums[i])
                
            sub.append([nums[i]])    
            all_subsets.extend(sub)
        
        return all_subsets
        