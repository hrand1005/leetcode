"""
class Solution:
    def arraySign(self, nums: List[int]) -> int:
        neg = 0
        for n in nums:
            if n == 0:
                return 0
            elif n < 0:
                neg += 1
        
        return 1 if neg % 2 == 0 else -1

from functools import reduce

class Solution:
    def arraySign(self, nums: List[int]) -> int:
        n = reduce(lambda x, y: x * y, nums)
        if n == 0:
            return 0
        elif n > 0:
            return 1
        return -1
        
"""

class Solution:
    def arraySign(self, nums: List[int]) -> int:
        if 0 in nums: return 0
        return 1 if len(list(filter(lambda x: x < 0, nums))) % 2 == 0 else -1