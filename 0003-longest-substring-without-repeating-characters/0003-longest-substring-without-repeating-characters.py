"""
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_sub = 0
        for i in range(len(s)):
            sub = ""
            occ = {}
            for j in range(i, len(s)):
                if occ.get(s[j]):
                    break
                else:
                    sub += s[j]
                    occ[s[j]] = 1

                if j == len(s) - 1:
                    return max(len(sub), max_sub)

            max_sub = max(len(sub), max_sub)

        return max_sub

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_sub = 0
        for i in range(len(s)):
            sub = ""
            for j in range(i, len(s)):
                if s[j] in sub:
                    break
                else:
                    sub += s[j]

            max_sub = max(len(sub), max_sub)       

        return max_sub    
""" 

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        last = {}
        left = 0
        max_sub = 0

        for i in range(len(s)):
            if last.get(s[i], -1) < left:
                # either there's no duplicate, or the duplicate is irrelevant
                # because it occurs at a place less than the left index
                max_sub = max(i - left + 1, max_sub)    
            else:
                # maintain a left index greater than the last duplicate character
                left = last[s[i]] + 1

            last[s[i]] = i

        return max_sub    



