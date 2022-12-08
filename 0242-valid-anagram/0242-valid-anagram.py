"""
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        occurrences = {}
        for char in s:
            if not occurrences.get(char):
                occurrences[char] = 1
            else:    
                occurrences[char] += 1
        
        for char in t:
            if occurrences.get(char, 0) == 0:
                return False
            occurrences[char] -= 1

        for key, value in occurrences.items():
            if value != 0:
                return False

        return True    
"""

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_list, t_list = list(s), list(t)
        s_list.sort()
        t_list.sort()
        return s_list == t_list
