class Solution:
    def hammingWeight(self, n: int) -> int:
        if n == 0:
            return 0
        
        count = 0
        while n > 1:
            if n % 2 == 1:
                count += 1
            n = n >> 1
            
        return count + 1  