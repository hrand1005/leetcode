"""
class Solution:
    def longestPalindrome(self, s: str) -> str:
        max_pal = s[0]

        for i in range(1, len(s)):
            odd = self.biggestPalindrome(i, i, s)
            even = self.biggestPalindrome(i-1, i, s)
            max_pal = max([odd, even, max_pal], key=len)

        return max_pal        

    def biggestPalindrome(self, left: int, right: int, s: str):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            right += 1
            left -= 1

        return s[left+1:right]    

    # Recursive...
    # def biggestPalindrome(self, left: int, right: int, s: str) -> str:
    #     if s[left] != s[right]:
    #         return s[left+1:right]
        
    #     if left == 0 or right == len(s) - 1:
    #         return s[left:right+1]

    #     return self.biggestPalindrome(left - 1, right + 1, s)
"""

class Solution:
    def longestPalindrome(self, s: str) -> str:
        table = [ [False] * len(s) for i in range(len(s)) ]
        for i in range(len(s)):
            table[i][i] = True

        max_pal = s[0]    

        for i in range(len(s)-1, -1, -1):
            for j in range(i+1, len(s)):
                if s[i] == s[j]:
                    if j - i < 2 or table[i+1][j-1]:
                        table[i][j] = True
                        max_pal = max([s[i:j+1], max_pal], key=len)

        return max_pal                


