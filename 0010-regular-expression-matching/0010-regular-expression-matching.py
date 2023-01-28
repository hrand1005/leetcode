"""
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        table = [[False] * (len(p)+1) for _ in range(len(s)+1)]
        table[0][0] = True
        
        for j in range(2, len(p)+1):
            if p[j-1] == "*":
                table[0][j] = table[0][j-2]
        
        for i in range(1, len(s)+1):
            for j in range(1, len(p)+1):
                if s[i-1] == p[j-1] or p[j-1] == ".":
                    table[i][j] = table[i-1][j-1]
                elif p[j-1] == "*":
                    empty = table[i][j-2]
                    nonempty = table[i-1][j] and (s[i-1] == p[j-2] or p[j-2] == ".")
                    table[i][j] = empty or nonempty
                    
        return table[-1][-1]
"""

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        self.cache = {
            ("", ""): True,
        }
        return self.is_match(s, p)
    
    def is_match(self, s: str, p: str) -> bool:
        if self.cache.get((s, p)) != None:
            return self.cache[(s, p)]
        if p == "":
            return False
        
        match = False
        if p[-1] == "*":
            idx = len(s)-1
            match = match or self.is_match(s, p[:-2])
            while 0 <= idx and (s[idx] == p[-2] or p[-2] == "."):
                match = match or self.is_match(s[:idx], p[:-2])
                idx -= 1
        elif s != "" and (s[-1] == p[-1] or p[-1] == "."):
            match = match or self.is_match(s[:-1], p[:-1])
            
        self.cache[(s, p)] = match
        return match
