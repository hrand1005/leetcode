"""
class Solution:
    def reverseBits(self, n: int) -> int:
        # convert to string, remove 0b from beginning
        s = bin(n)[2:]
        # reverse string
        s = s[::-1]
        # pad with zeroes up to 32 to the end -- remember, these
        # zeroes occurred at the beginning in the initial integer
        s = s + "0"*(32-len(s))
        # convert back to integer
        s = int(s, 2)
        return s
"""

class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        for i in range(32):
            res = res << 1 # left shift res
            res += n%2 # adds 1 if n's least significant bit is 1
            n = n >> 1 # right shift n
        
        return res