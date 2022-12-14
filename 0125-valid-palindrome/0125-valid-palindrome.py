"""
VALID_CHARS = 'abcdefghijklmnopqrstuvwxyz0123456789'

class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = self.sanitizeInput(s)
        
        for i in range(len(s)):
            if s[i] != s[len(s)-1-i]:
                return False
        
        return True
    
    def sanitizeInput(self, s: str) -> str:
        s = s.lower()
        
        i = 0
        while i < len(s):
            if s[i] not in VALID_CHARS:
                char = s[i]
                s = s[0:i] + s[i+1:]
            else:
                i += 1
        
        return s
"""
class Solution:
    def isPalindrome(self, s: str) -> bool:
        i = 0
        j = len(s) - 1
        while i < j:    
            while i < j and not s[i].isalnum():
                i += 1
                
            while i < j and not s[j].isalnum():
                j -= 1
            
            if i >= j:
                return True
            
            if s[i].lower() != s[j].lower():
                return False
                
            i += 1
            j -= 1

        return True       