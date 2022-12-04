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