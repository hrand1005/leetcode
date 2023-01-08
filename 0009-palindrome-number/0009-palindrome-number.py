"""
class Solution:
    def isPalindrome(self, x: int) -> bool:
        x_str = str(x)
        return self.check_palindrome(x_str)
    
    def check_palindrome(self, s: str) -> bool:
        if len(s) <= 1:
            return True
        if s[0] != s[-1]:
            return False
        return self.check_palindrome(s[1:-1])
            
"""

class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        before = x
        after = 0
        while x > 0:
            after *= 10
            after += (x % 10)
            x //= 10
        return before == after    