class Solution:
    def myPow(self, x: float, n: int) -> float:
        if x == 0: 
            return 0
        
        negative_exponent = False
        if n < 0:
            n = -n
            negative_exponent = True
        
        res = 1
        while n > 0:
            scale = 1
            new_x = x
            while scale * 2 < n:
                new_x *= new_x
                scale *= 2

            n -= scale
            res *= new_x
                
        if negative_exponent:
            return 1 / res
        
        return res
        