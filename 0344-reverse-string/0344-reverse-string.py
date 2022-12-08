"""
class Solution:
    def reverseString(self, s: List[str]) -> None:
        for i in range(len(s)//2):
            j = len(s) - 1 - i
            s[i], s[j] = s[j], s[i]

class Solution:
    def reverseString(self, s: List[str]) -> None:
        s.reverse()
""" 

class Solution:
    def reverseString(self, s: List[str]) -> None:
        s[:] = s[::-1]