# brute force 10 min
"""
class Solution:
    def maxArea(self, height: List[int]) -> int:
        max_water = 0
        for i in range(0, len(height)):
            for j in range(1, len(height)):
                length = j - i
                depth = min(height[i], height[j])
                water = length * depth
                max_water = max(water, max_water)
                    
        return max_water 
"""    

# 32 minutes
"""
class Solution:
    def maxArea(self, height: List[int]) -> int:
        lefts, rights = [0] * len(height), [0] * len(height)
        
        lefts[0] = height[0]
        for i in range(1, len(height)):
            if height[i] > lefts[len(lefts)-1]:
                lefts[i] = height[i]
            else:
                lefts[i] = 0
        
        rights[-1] = height[-1]
        for i in range(len(height)-1, -1, -1):
            if height[i] > rights[-1]:
                rights[i] = height[i]
            else:
                rights[i] = 0
        
        max_water = 0
        for i in range(len(lefts)):
            for j in range(i+1, len(height)):
                length = j - i
                depth = min(height[i], height[j])
                water = length * depth
                max_water = max(water, max_water)
        
        return max_water
"""

class Solution:
    def maxArea(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        max_water = 0
        
        while l < r:
            depth = min(height[l], height[r])
            length = r - l
            max_water = max(max_water, length * depth)
            
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        
        return max_water