"""
class Solution(object):
    def isPalindrome(self, s):
        if len(s) <= 1:
            return True
        first, last = s[0], s[-1]
        if first.isalnum() and last.isalnum():
            if lower(first) != lower(last):
                return False
            return self.isPalindrome(s[1:-1])
        elif not first.isalnum() and not last.isalnum():
            return self.isPalindrome(s[1:-1])
        elif not first.isalnum():
            return self.isPalindrome(s[1:])
        else:
            return self.isPalindrome(s[:-1])
"""

class Solution(object):
    def isPalindrome(self, s):
        l, r = 0, len(s)-1
        while l < r:
            while l < r and not s[l].isalnum():
                l += 1
            while r > l and not s[r].isalnum():
                r -= 1
            if l < r and s[l].lower() != s[r].lower():
                return False
            l += 1
            r -= 1
        return True