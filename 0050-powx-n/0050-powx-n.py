"""
class Solution:
    def myPow(self, x: float, n: int) -> float:
        if x == 0: 
            return 0
        
        neg_exp = False
        if n < 0:
            n = -n
            neg_exp = True
        
        res = 1
        while n > 0:
            scale = 1
            new_x = x
            while scale * 2 < n:
                new_x *= new_x
                scale *= 2

            n -= scale
            res *= new_x
                
        if neg_exp:
            return 1 / res
        
        return res
"""

class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        if n < 0:
            return 1 / self.myPow(x, -n)
        if n % 2:
            return x * self.myPow(x, n-1)
        
        return self.myPow(x*x, n/2)