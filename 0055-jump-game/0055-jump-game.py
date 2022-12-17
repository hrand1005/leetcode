"""
class Solution:
    def canJump(self, nums: List[int]) -> bool:
        stack = [(0, nums[0])]
        target = len(nums) - 1
        visited = {}
        
        while stack:
            idx, val = stack.pop()
            if idx == target:
                return True
            
            if visited.get(idx) != None:
                continue
            
            visited[idx] = True
            for i in range(val, 0, -1):
                if idx + i > len(nums) - 1:
                    return True
                if visited.get(idx+i) == None:
                    stack.append((idx+i, nums[idx+i]))
            
        return False
        
"""

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        target = len(nums) - 1
        max_reach = nums[0]
        
        for i in range(1, len(nums)):
            if max_reach < i:
                return False
            
            max_reach = max(max_reach, i + nums[i])
            if max_reach >= target:
                return True
        
        return max_reach >= target
             
            