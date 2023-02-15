class Solution(object):
    def isHappy(self, n):
        seen = {}
        dsum = self.digit_sum(n)
        seen[dsum] = True
        while dsum != 1:
            dsum = self.digit_sum(dsum)
            if seen.get(dsum, False):
                return False
            seen[dsum] = True
        return True    
    
    def digit_sum(self, num):
        total = 0
        for c in str(num):
            total += int(c) ** 2
        return total    
        