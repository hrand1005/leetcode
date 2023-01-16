class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        last = len(s)-1
        while 0 <= last and s[last] == " ":
            last -= 1
            
        count = 0    
        while 0 <= last and s[last] != " ":
            count += 1
            last -= 1
        
        return count