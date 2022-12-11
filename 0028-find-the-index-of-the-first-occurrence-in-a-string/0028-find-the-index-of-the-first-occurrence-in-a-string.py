class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        for i in range(len(haystack)):
            if haystack[i] == needle[0]:
                if i + len(needle) > len(haystack):
                    return -1
                check = haystack[i:i+len(needle)]
                if check == needle:
                    return i
        
        return -1