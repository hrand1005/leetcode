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
            
        