class Solution:
    def getSum(self, a: int, b: int) -> int:
        carry = 0
        mask = 0xffffffff
        while b & mask != 0:
            a, b = a ^ b, (a & b) << 1
            
        if mask < b:
            return a & mask
        return a