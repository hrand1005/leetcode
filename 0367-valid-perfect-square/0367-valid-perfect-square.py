class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        i = 1
        while i**2 < num:
            n = i
            while n**2 < num:
                i = n
                n *= 2
            i += 1    
            
        return i**2 == num
        