"""
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

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        all_subsets = []
        self.dfs(nums, [], all_subsets)
        return all_subsets
    
    def dfs(self, nums: List[int], current: List[int], all_subsets: List[List[int]]):
        all_subsets.append(current)
        for i in range(len(nums)):
            self.dfs(nums[i+1:], current+[nums[i]], all_subsets)
        
"""

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        all_subsets = [[]]
        for n in nums:
            all_subsets += [s + [n] for s in all_subsets]
        
        return all_subsets
