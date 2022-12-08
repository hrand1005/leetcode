"""
class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n < 1:
            return False

        while n != 1:
            if n % 3 != 0:
                return False
    
            n = n // 3    

        return True

class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n < 1:
            return False
        elif n == 1:
            return True    

        if n % 3 != 0:
            return False    

        return self.isPowerOfThree(n//3)    
""" 

from math import log

class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n < 1:
            return False

        x = round(log(n, 3)) # use round in case imprecise result, e.g. 2.99999999999   
        return n == (3 ** x)