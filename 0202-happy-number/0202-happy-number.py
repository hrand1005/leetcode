"""
class Solution:
    def isHappy(self, n: int) -> bool:
        occurred = {}
        
        while occurred.get(n) == None:
            occurred[n] = True
            num_string = str(n)
            this_sum = 0
            
            for digit in num_string:
                this_sum += int(digit) ** 2
        
            if this_sum == 1:
                return True
            
            n = this_sum
        
        return False
"""

class Solution:
    def isHappy(self, n: int) -> bool:
        slow = n
        fast = self.sum_squared_digits(slow)
        
        while slow != fast and fast != 1:
            slow = self.sum_squared_digits(slow)
            fast = self.sum_squared_digits(fast)
            
            # not really necessary to check if fast == 1
            # because it will remain 1 in perpetuity
            if fast == slow:
                return False
            
            # go again for fast!
            fast = self.sum_squared_digits(fast)
            
        if fast == 1:
            return True
        
        return False    
        
    def sum_squared_digits(self, n: int) -> int:
        digit_sum = 0
        n_string = str(n)
        for digit in n_string:
            digit_sum += int(digit) ** 2
            
        return digit_sum    