class Solution:
    def reverseBits(self, n: int) -> int:
        # convert to string, remove 0b from beginning
        s = bin(n)[2:]
        # reverse string
        s = s[::-1]
        # pad with zeroes up to 32 to the end -- remember, 
        # these zeroes occurred at the beginning in the initial
        # integer
        s = s + "0"*(32-len(s))
        # convert back to integer
        s = int(s, 2)
        return s